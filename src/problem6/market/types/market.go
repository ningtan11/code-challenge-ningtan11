package types

import (
	"reflect"
	"regexp"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// System-wide maximum market lotsize and ticksize
var (
	MaxTickSize   = sdk.MustNewDecFromStr("1000000000000000000000000000000") // 30 zeros
	MaxLotSize, _ = sdk.NewIntFromString("1000000000000000000000000000000")  // 30 zeros
)

// MarketType to string literal map
const (
	SpotMarket    string = "spot"
	FuturesMarket string = "futures"
)

const (
	MaxMarketNameLength        = 128
	MaxMarketDisplayNameLength = 128
)

// Mutate changes fields in Market from MarketParams that are non-nil.
func (m *Market) Mutate(mm MarketParams) error {
	mVal := reflect.ValueOf(m).Elem()
	mmVal := reflect.ValueOf(&mm).Elem()

	for i := 0; i < mmVal.NumField(); i++ {
		mmFieldValue := mmVal.Field(i)            // get mm field value
		mmFieldName := mmVal.Type().Field(i).Name // get mm field name
		if mmFieldName == "Name" {
			// ignore the ID field
			continue
		}
		if !mmFieldValue.IsNil() {
			// field is set, mutate the value
			mField := mVal.FieldByName(mmFieldName) // get m field by mm's field name
			mField.Set(mmFieldValue.Elem())         // set m field to mm's value
		}
	}

	return nil
}

// PrecisionDifference returns the difference in
// the precision of the quote currency and
// the precision of the base currency.
func (m Market) PrecisionDifference() int64 {
	return m.QuotePrecision - m.BasePrecision
}

// GetStoreKey returns the store key for the market which is just the market name in bytes
func (m Market) GetStoreKey() []byte {
	return []byte(m.Name)
}

// IsMarketActive checks if market is active
func (m Market) IsMarketActive(ctx sdk.Context) bool {
	return m.IsActive && m.IsMarketBeforeExpiry(ctx)
}

// IsMarketBeforeExpiry checks if market is before expiry
func (m Market) IsMarketBeforeExpiry(ctx sdk.Context) bool {
	if m.MarketType != FuturesMarket {
		return true
	}
	if m.IsPerpetualFutures() {
		return true
	}
	return ctx.BlockTime().Before(m.ExpiryTime) // edit here
}

// FilterMarketsBeforeExpiry filters markets before expiry
func FilterMarketsBeforeExpiry(ctx sdk.Context, markets []Market) []Market {
	var res []Market
	for _, market := range markets {
		if market.IsMarketBeforeExpiry(ctx) {
			res = append(res, market)
		}
	}
	return res
}

// FilterMarketsToSettle filters markets to settle
func FilterMarketsToSettle(ctx sdk.Context, markets []Market) []Market {
	var marketsToSettle []Market
	for _, market := range markets {
		if !market.IsMarketBeforeExpiry(ctx) && !market.IsSettled {
			marketsToSettle = append(marketsToSettle, market)
		}
	}
	return marketsToSettle
}

// IsInMarkets returns true if this market is found in markets, matching by name
func (m Market) IsInMarkets(markets []Market) bool {
	for _, market := range markets {
		if market.Name == m.Name {
			return true
		}
	}
	return false
}

// IsFutures checks if market is futures
func (m Market) IsFutures() bool {
	return m.MarketType == FuturesMarket
}

func (m Market) IsSpot() bool {
	return m.MarketType == SpotMarket
}

func (m Market) IsPerpetualFutures() bool {
	return m.IsFutures() && m.ExpiryTime.Unix() == 0
}

// MaxLeverageDec returns the max leverage setting for the market.
// This ignores users' positions and returns just the max leverage
// for the base initial margin requirement where
// MaxLeverage = 1 / InitialMarginBase.
func (m Market) MaxLeverageDec() sdk.Dec {
	return sdk.OneDec().QuoTruncate(m.InitialMarginBase)
}

// RequiredMaintenanceMargin returns the required maintenance margin
// ratio for the given position size.
// Formula: MM = IM * MMR
func (m Market) RequiredMaintenanceMargin(size sdkmath.Int) sdk.Dec {
	return m.RequiredInitialMargin(size).Mul(m.MaintenanceMarginRatio)
}

// RequiredInitialMargin returns the required maintenance margin
// ratio for the given position size.
// Formula: IM = IMB + floor(abs(size) / RSS) * IMS
func (m Market) RequiredInitialMargin(size sdkmath.Int) sdk.Dec {
	if size.IsNegative() {
		size = size.MulRaw(-1)
	}
	riskStep := size.Quo(m.RiskStepSize)
	return m.InitialMarginBase.Add(m.InitialMarginStep.MulInt(riskStep))
}

// FeeMargin returns the margin required for fees. This is derived from
// the maximum potential entry fee at the given qty and price.
// The isWaiting flag should be set to true if the fees is to be estimated
// for a new order (that is not yet on the order book).
//
// This method only considers the positive part of fees and fees that are negative
// (a rebate) are ignored and treated as zero.
//
// This method is only supported for futures market as fee margin spot markets is
// always zero. The fee for spot market is deducted from the taken amount and
// additional margin is never required.
func (m Market) FeeMargin(quantity sdkmath.Int, price sdk.Dec, isWaiting bool, marketFees TradingFees) sdk.Coin {
	if m.MarketType != FuturesMarket {
		panic("FeeMargin called for non-futures market!")
	}

	// always assume max fee since edited order might be taker/maker which needs to pay for fees
	maxFee := sdk.MaxDec(marketFees.TakerFee, marketFees.MakerFee)
	amount := sdk.NewDecFromInt(quantity).Mul(maxFee)

	return sdk.NewCoin(m.Quote, amount.Mul(price).Ceil().TruncateInt()) // round up
}

// MaxSizeGivenStepSize returns the maximum permitted size for a given step size
// example: if every step is 10 btc, step 0 should return 9.999 btc
func (m Market) MaxSizeGivenStepSize(stepSize sdkmath.Int) sdkmath.Int {
	return (stepSize.Add(sdk.OneInt()).Mul(m.RiskStepSize)).Sub(m.LotSize)
}

// MinSizeGivenStepSize returns the minimum permitted size for a given step size
// example: if every step is 10 btc, step 0 should return 0.001 btc
func (m Market) MinSizeGivenStepSize(stepSize sdkmath.Int) sdkmath.Int {
	return stepSize.Mul(m.RiskStepSize).Add(m.LotSize)
}

// // MarketEvent prepares Market to be emitted
// func (m Market) MarketEvent(t transitiontypes.StateTransition) MarketEvent {
//	return MarketEvent{
//		Market: &m,
//		Type:   string(t),
//	}
// }

func (m Market) ValidateBasic() error {
	// general validations
	switch {
	case m.Name == "":
		return errorsmod.Wrap(ErrInvalidMarket, "name must not be empty")
	case !regexp.MustCompile(`^[a-zA-Z0-9_./]+$`).MatchString(m.Name):
		return errorsmod.Wrap(ErrInvalidMarket, "name must contain only a-z, A-Z, 0-9, _, '.' or '/'")
	case len(m.Name) > MaxMarketNameLength:
		return errorsmod.Wrapf(ErrInvalidMarket, "name must be equal or less than %v in length", MaxMarketNameLength)
	case m.DisplayName == "":
		return errorsmod.Wrap(ErrInvalidMarket, "display_name must not be empty")
	case !regexp.MustCompile(`^[a-zA-Z0-9_. ]+$`).MatchString(m.DisplayName):
		return errorsmod.Wrap(ErrInvalidMarket, "display_name must contain only a-z, A-Z, 0-9, _, space, or '.'")
	case len(m.DisplayName) > MaxMarketDisplayNameLength:
		return errorsmod.Wrapf(ErrInvalidMarket, "display_name must be equal or less than %v in length", MaxMarketDisplayNameLength)
	case m.Description == "":
		return errorsmod.Wrap(ErrInvalidMarket, "description must not be empty")
	case m.Base == "":
		return errorsmod.Wrap(ErrInvalidMarket, "base must not be empty")
	case m.Quote == "":
		return errorsmod.Wrap(ErrInvalidMarket, "base must not be empty")
	case !m.LotSize.IsPositive():
		return errorsmod.Wrap(ErrInvalidMarket, "lot_size must be more than zero")
	case m.LotSize.GT(MaxLotSize):
		return errorsmod.Wrap(ErrInvalidMarket, "lot_size is too large")
	case !m.TickSize.IsPositive():
		return errorsmod.Wrap(ErrInvalidMarket, "tick_size must be positive")
	case !m.TickSize.MulInt(m.LotSize).IsInteger():
		return errorsmod.Wrap(ErrInvalidMarket, "tick_size * lot_size must be an integer")
	case m.TickSize.GT(MaxTickSize):
		return errorsmod.Wrap(ErrInvalidMarket, "tick_size is too large")
	case !m.MinQuantity.IsPositive():
		return errorsmod.Wrap(ErrInvalidMarket, "min_quantity must be positive")
	case m.MinQuantity.LT(m.LotSize):
		return errorsmod.Wrap(ErrInvalidMarket, "min_quantity must be greater or equal to lot_size")
	case !m.MinQuantity.Mod(m.LotSize).IsZero():
		return errorsmod.Wrapf(ErrInvalidMarket, "MinQuantity: %v is not divisible by lotSize: %v", m.MinQuantity, m.LotSize)
	}

	// validations based on MarketType
	switch m.MarketType {
	case SpotMarket:
		switch {
		case !m.RiskStepSize.IsZero():
			return errorsmod.Wrap(ErrInvalidMarket, "risk_step_size for spot markets must be zero")
		case !m.InitialMarginBase.Equal(sdk.OneDec()):
			return errorsmod.Wrap(ErrInvalidMarket, "initial_margin_base for spot markets must be 100%")
		case !m.InitialMarginStep.IsZero():
			return errorsmod.Wrap(ErrInvalidMarket, "initial_margin_step for spot markets must be zero")
		case !m.MaintenanceMarginRatio.IsZero():
			return errorsmod.Wrap(ErrInvalidMarket, "maintenance_margin_ratio for spot markets must be zero")
		case !m.MaxLiquidationOrderTicket.IsZero():
			return errorsmod.Wrap(ErrInvalidMarket, "max_liquidation_order_ticket for spot markets must be zero")
		case !m.ImpactSize.IsZero():
			return errorsmod.Wrap(ErrInvalidMarket, "impact_size for spot markets must be zero")
		case m.MarkPriceBand != 0:
			return errorsmod.Wrap(ErrInvalidMarket, "mark_price_band for spot markets must be zero")
		case m.LastPriceProtectedBand != 0:
			return errorsmod.Wrap(ErrInvalidMarket, "last_price_protected_band for spot markets must be zero")
		case m.IndexOracleId != "":
			return errorsmod.Wrap(ErrInvalidMarket, "index_oracle_id for spot markets must be empty")
		case m.ExpiryTime.Unix() != 0:
			return errorsmod.Wrap(ErrInvalidMarket, "expiry_time for spot markets must be zero")
		case m.MaxLiquidationOrderDuration.Nanoseconds() != 0:
			return errorsmod.Wrap(ErrInvalidMarket, "max_liquidation_order_duration for spot markets must be zero")
		}
	case FuturesMarket:
		switch {
		case !m.InitialMarginBase.IsPositive():
			return errorsmod.Wrap(ErrInvalidMarket, "initial_margin_base for future markets must be positive")
		case !m.MaintenanceMarginRatio.IsPositive() || m.MaintenanceMarginRatio.GTE(sdk.OneDec()):
			return errorsmod.Wrap(ErrInvalidMarket, "maintenance_margin_ratio for futures markets must between zero and one")
		case m.RiskStepSize.IsNegative():
			return errorsmod.Wrap(ErrInvalidMarket, "risk_step_size must not be negative")
		case m.InitialMarginStep.IsNegative():
			return errorsmod.Wrap(ErrInvalidMarket, "initial_margin_step must not be negative")
		case m.RiskStepSize.IsZero() != m.InitialMarginStep.IsZero():
			return errorsmod.Wrap(ErrInvalidMarket, "risk_step_size and initial_margin_step must be either both zero or both not zero")
		case m.MaxLiquidationOrderTicket.MulRaw(1000).LT(m.MinQuantity):
			return errorsmod.Wrap(ErrInvalidMarket, "max_liquidation_order_ticket must be at least 1000x of min_quantity")
		case !m.ImpactSize.IsPositive():
			return errorsmod.Wrap(ErrInvalidMarket, "impact_size for futures markets must be positive")
		case m.MarkPriceBand == 0 || m.MarkPriceBand > 20000:
			return errorsmod.Wrap(ErrInvalidMarket, "mark_price_band for futures markets must be between 0 and 20000")
		case m.LastPriceProtectedBand == 0 || m.LastPriceProtectedBand > 20000:
			return errorsmod.Wrap(ErrInvalidMarket, "last_price_protected_band must be between 0 and 20000")
		case m.IndexOracleId == "":
			return errorsmod.Wrap(ErrInvalidMarket, "index_oracle_id (empty) is required for futures markets")
		case m.MaxLiquidationOrderDuration.Seconds() < 30:
			return errorsmod.Wrap(ErrInvalidMarket, "max_liquidation_order_duration must be at least 30 seconds")

		case m.ExpiryTime.Unix() >= 253402300800:
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "market expiry time's unix must be less than 253402300800")
		}
	default:
		return errorsmod.Wrap(ErrInvalidMarket, "market_type is invalid")
	}
	return nil
}
