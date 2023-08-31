package types

import (
	"regexp"

	errorsmod "cosmossdk.io/errors"
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
