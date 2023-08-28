package keeper

import (
	"fmt"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	transitiontypes "github.com/Switcheo/carbon/common/types"
	"github.com/Switcheo/carbon/common/utils"
	cointypes "github.com/Switcheo/carbon/x/coin/types"
	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ====================== HANDLER FUNCTION FOR MARKET CREATION
func (k Keeper) CreateMarket(ctx sdk.Context, marketType string, base string, quote string,
	currentBasePrice sdk.Dec, currentQuotePrice sdk.Dec, indexOracleId string, expiryTime time.Time) (market types.Market, err error) {
	params := k.GetParams(ctx)

	if base == quote {
		return market, errorsmod.Wrapf(types.ErrSameBaseAndQuoteToken, "token: %s", base)
	}

	market, err = k.generateMarketAttributes(ctx, params, marketType, base, quote,
		currentBasePrice, currentQuotePrice, indexOracleId, expiryTime)
	if err != nil {
		return market, err
	}

	if err = market.ValidateBasic(); err != nil {
		return market, err
	}

	if isUnique := k.IsUniqueMarket(ctx, market.MarketType, market.Base, market.Quote, market.ExpiryTime); !isUnique {
		return market, errorsmod.Wrapf(types.ErrMarketNotUnique, "MarketType: %s, BaseDenom: %s, QuoteDenom: %s, ExpiryTime: %s", market.MarketType, market.Base, market.Quote, market.ExpiryTime)
	}

	if found := k.HasMarket(ctx, market.Name); found {
		return market, errorsmod.Wrap(types.ErrDuplicateName, market.Name)
	}

	if k.IsDisplayNameUsed(ctx, market.DisplayName) {
		return market, errorsmod.Wrap(types.ErrDuplicateDisplayName, market.DisplayName)
	}

	if err = k.ValidateMarket(ctx, market); err != nil {
		return market, err
	}

	if err = k.setTokenPrecision(ctx, &market); err != nil {
		panic(err)
	}

	k.SetCreatedMarket(ctx, market)

	return market, err
}

// ====================== GETTER AND SETTERS FOR THE KV STORE
// ====================== Question
// GetAllMarket returns all markets from the store
func (k Keeper) GetAllMarket(ctx sdk.Context) (list []types.Market, err error) {
	// Implement a new method to obtain all markets within the store
}

// SetCreatedMarket set a specific market in the store
func (k Keeper) SetCreatedMarket(ctx sdk.Context, market types.Market) {
	store := k.Store(ctx, types.MarketKey)
	market.ExpiryTime = market.ExpiryTime.UTC() // stardardise ExpiryTime, used in validate unique market
	b := k.Marshal(&market)
	store.Set([]byte(market.Name), b)
}

// HasMarket checks if the market exists in the store
func (k Keeper) HasMarket(ctx sdk.Context, name string) bool {
	store := k.Store(ctx, types.MarketKey)
	return store.Has([]byte(name))
}

// RemoveMarket removes a market from the store
func (k Keeper) RemoveMarket(ctx sdk.Context, name string) {
	store := k.Store(ctx, types.MarketKey)
	store.Delete([]byte(name))
}

// GetMarket returns a market from its name, and true or false depending
// on if the market was found.
func (k Keeper) GetMarket(ctx sdk.Context, name string) (market types.Market, found bool) {
	store := k.Store(ctx, types.MarketKey)
	bz := store.Get([]byte(name))
	if bz == nil {
		return market, found
	}
	k.Unmarshal(bz, &market)
	return market, true
}

// ====================== HELPER FUNCTION FOR GENERATING MARKET
func (k Keeper) generateMarketAttributes(ctx sdk.Context, params types.Params, marketType string, base string, quote string,
	currentBasePrice sdk.Dec, currentQuotePrice sdk.Dec, indexOracleId string, expiryTime time.Time) (m types.Market, err error) {
	baseToken, found := k.ck.GetToken(ctx, base)
	if !found {
		return m, errorsmod.Wrap(cointypes.ErrTokenNotFoundType, base)
	}
	quoteToken, found := k.ck.GetToken(ctx, quote)
	if !found {
		return m, errorsmod.Wrap(cointypes.ErrTokenNotFoundType, quote)
	}
	if !baseToken.IsActive {
		return m, errorsmod.Wrap(cointypes.ErrTokenInactive, base)
	}
	if !quoteToken.IsActive {
		return m, errorsmod.Wrap(cointypes.ErrTokenInactive, quote)
	}
	lotSize, errLotSize := generateLotSize(params, currentBasePrice, baseToken.Decimals)
	if errLotSize != nil {
		return m, errLotSize
	}
	tickSize, errTickSize := generateTickSize(params, currentBasePrice, currentQuotePrice, baseToken.Decimals, quoteToken.Decimals, lotSize)
	if errTickSize != nil {
		return m, errTickSize
	}
	riskStepSize := sdk.ZeroInt()
	initialMarginBase := sdk.OneDec()
	initialMarginStep := sdk.ZeroDec()
	maintenanceMarginRatio := sdk.ZeroDec()
	maxLiquidationOrderTicket := sdk.ZeroInt()
	impactSize := sdk.ZeroInt()
	markPriceBand := uint32(0)
	lastPriceProtectedBand := uint32(0)
	maxLiquidationOrderDuration := time.Duration(0)
	tradingBandwidth := uint32(0)
	createdBlockHeight := uint64(ctx.BlockHeight())

	marketNameSequence := k.GetMarketNameSequence(ctx)

	displayName := generateDisplayName(marketType, baseToken.Symbol, quoteToken.Symbol, expiryTime)
	if k.IsDisplayNameUsed(ctx, displayName) {
		// append the market sequence to allow tokens with same symbol to be created:
		displayName = fmt.Sprintf("%s.%v", displayName, marketNameSequence)
	}

	m = types.Market{
		Name:                        fmt.Sprintf("cmkt/%v", marketNameSequence),
		DisplayName:                 displayName,
		Description:                 generateDescription(marketType, baseToken.Name, baseToken.Symbol, quoteToken.Symbol, expiryTime),
		MarketType:                  marketType,
		Base:                        baseToken.Denom,
		Quote:                       quoteToken.Denom,
		BasePrecision:               baseToken.Decimals,
		QuotePrecision:              quoteToken.Decimals,
		LotSize:                     lotSize,
		TickSize:                    tickSize,
		MinQuantity:                 generateMinQuantity(params, currentBasePrice, baseToken.Decimals, lotSize),
		RiskStepSize:                riskStepSize,
		InitialMarginBase:           initialMarginBase,
		InitialMarginStep:           initialMarginStep,
		MaintenanceMarginRatio:      maintenanceMarginRatio,
		MaxLiquidationOrderTicket:   maxLiquidationOrderTicket,
		MaxLiquidationOrderDuration: maxLiquidationOrderDuration,
		ImpactSize:                  impactSize,
		MarkPriceBand:               markPriceBand,
		LastPriceProtectedBand:      lastPriceProtectedBand,
		IndexOracleId:               indexOracleId,
		ExpiryTime:                  expiryTime.UTC(),
		IsActive:                    true,
		TradingBandwidth:            tradingBandwidth,
		CreatedBlockHeight:          createdBlockHeight,
	}
	return m, err
}

func generateDisplayName(marketType string, base string, quote string, expiryTime time.Time) string {
	if isFutures(marketType) {
		if expiryTime.Unix() == 0 {
			return strings.ToUpper(base + "_PERP." + quote)
		}
		return strings.ToUpper(base + "_" + expiryTime.Format("02Jan06") + "." + quote)
	} else {
		return strings.ToUpper(base + "_" + quote)
	}
}

func generateDescription(marketType string, baseName string, base string, quote string, expiryTime time.Time) string {
	if isFutures(marketType) {
		if expiryTime.Unix() == 0 {
			return baseName + " Perpetual Futures"
		}
		return baseName + " Futures Expiring " + expiryTime.Format("2 January 2006")
	} else {
		return strings.ToUpper(base+"/"+quote) + " Spot Market"
	}
}

// Number of lots that will be worth DefaultLotSizeUsd, which is usually ~$0.10
// Lot size is rounded to 1 s.f. and then the first digit is moved to the nearest 1/10.
// e.g. 51234567 => 50000000 => 100000000
func generateLotSize(params types.Params, currentBasePrice sdk.Dec, decimals int64) (sdkmath.Int, error) {
	target := utils.DecShift(params.DefaultLotSizeUsd.Quo(currentBasePrice), decimals).Ceil().TruncateInt()
	if target.LT(sdk.OneInt()) {
		return sdk.ZeroInt(), errorsmod.Wrapf(types.ErrInvalidLotSize, "lot size: %v, please recheck current base price", target)
	}
	digits := getDigits(target)
	factor := sdkmath.NewIntWithDecimal(1, digits-1)
	firstDigit := target.Quo(factor)
	return roundToNearestTen(firstDigit).Mul(factor), nil
}

// Price tick to match DefaultTickSizeUsd, which is usually ~$0.01
// Tick size is rounded to the nearest (1/lot size), then 1 s.f., then the first digit is moved to the nearest 1/5/10.
// e.g. 412.34567 => 412.30000 => 400.00000 => 500.00000
// e.g. 912.34567 => 912.30000 => 900.00000 => 1000.00000
func generateTickSize(params types.Params, currentBasePrice, currentQuotePrice sdk.Dec, baseDecimals, quoteDecimals int64, lotSize sdkmath.Int) (sdk.Dec, error) {
	pricePrecisionDiff := quoteDecimals - baseDecimals // price = quote/base
	tickSizeRatioOfPrice := params.DefaultTickSizeUsd
	target := utils.DecShift(tickSizeRatioOfPrice.Mul(currentBasePrice).Quo(currentQuotePrice), pricePrecisionDiff) // normalize the tick size to the relative token usd prices
	targetInt := target.MulInt(lotSize).Ceil().TruncateInt()                                                        // ensure divisibility (i.e. lot size * tick size % 1 = 0

	if targetInt.LT(sdk.OneInt()) {
		return sdk.ZeroDec(), errorsmod.Wrapf(types.ErrInvalidTickSize, "tick size: %v, please recheck current base price and current quote price", targetInt)
	}
	digits := getDigits(targetInt)
	factor := sdkmath.NewIntWithDecimal(1, digits-1)
	firstDigit := targetInt.Quo(factor)
	return sdk.NewDecFromInt(roundToNearestFive(firstDigit).Mul(factor)).QuoInt(lotSize), nil
}

// Number of lots that will be worth DefaultMinQuantityUsd, which is usually ~$1, rounded to 2 s.f.
func generateMinQuantity(params types.Params, currentBasePrice sdk.Dec, baseDecimals int64, lotSize sdkmath.Int) sdkmath.Int {
	target := roundToSignificantFigures(utils.DecShift(params.DefaultMinQuantityUsd.Quo(currentBasePrice), baseDecimals).Ceil().TruncateInt(), 2)
	multiple := sdk.NewDecFromInt(target).QuoRoundUp(sdk.NewDecFromInt(lotSize)).Ceil().TruncateInt()
	return multiple.Mul(lotSize)
}

// Number of lots that will be worth DefaultImpactSizeUsd, which is usually ~$100k, rounded to 2 s.f.
func generateImpactSize(params types.Params, currentBasePrice sdk.Dec, baseDecimals int64) sdkmath.Int {
	target := utils.DecShift(params.DefaultImpactSizeUsd.Quo(currentBasePrice), baseDecimals).Ceil().TruncateInt()
	return roundToSignificantFigures(target, 2)
}

func getDigits(i sdkmath.Int) (digits int) {
	if i.LT(sdk.OneInt()) {
		panic("getNumDigits cannot be used on int < 1")
	}
	for i.GTE(sdk.OneInt()) {
		i = i.QuoRaw(10)
		digits += 1
	}
	return
}

func roundToNearestTen(i sdkmath.Int) sdkmath.Int {
	if i.LT(sdk.OneInt()) {
		panic("roundToNearestTen cannot be used on int < 1")
	}
	if i.GT(sdk.NewInt(9)) {
		panic("roundToNearestTen cannot be used on int > 9")
	}
	if i.LT(sdk.NewInt(5)) {
		return sdk.OneInt()
	}
	return sdk.NewInt(10)
}

func roundToNearestFive(i sdkmath.Int) sdkmath.Int {
	if i.LT(sdk.OneInt()) {
		panic("roundToNearestFive cannot be used on int < 1")
	}
	if i.GT(sdk.NewInt(9)) {
		panic("roundToNearestFive cannot be used on int > 9")
	}
	if i.LT(sdk.NewInt(4)) {
		return sdk.OneInt() // 1, 2, 3
	}
	if i.LT(sdk.NewInt(7)) {
		return sdk.NewInt(5) // 4, 5, 6
	}
	return sdk.NewInt(10) // 7, 8, 9
}

func roundToSignificantFigures(i sdkmath.Int, sf int) sdkmath.Int {
	if sf < 1 {
		panic("roundToSignificantFigures needs a positive significant figure")
	}
	digits := getDigits(i)
	exponent := digits - sf
	if exponent < 0 {
		exponent = 0
	}
	factor := sdkmath.NewIntWithDecimal(1, exponent)
	return i.Quo(factor).Mul(factor)
}

func (k Keeper) setTokenPrecision(ctx sdk.Context, market *types.Market) error {
	// ensures base and quote tokens are already listed
	base, found := k.ck.GetToken(ctx, market.Base)
	if !found {
		return errorsmod.Wrap(cointypes.ErrTokenNotFoundType, market.Base)
	}
	quote, found := k.ck.GetToken(ctx, market.Quote)
	if !found {
		return errorsmod.Wrap(cointypes.ErrTokenNotFoundType, market.Quote)
	}
	market.BasePrecision = base.Decimals
	market.QuotePrecision = quote.Decimals
	return nil
}

// IsUniqueMarket checks if there exist a market with same base, quote. Futures with different expiry time are unique.
func (k Keeper) IsUniqueMarket(ctx sdk.Context, newMarketType, baseDenom, quoteDenom string, expiryTime time.Time) bool {
	for _, market := range k.GetAllMarket(ctx) {
		if market.MarketType == newMarketType && market.Base == baseDenom && market.Quote == quoteDenom && market.ExpiryTime.UTC() == expiryTime.UTC() {
			return false
		}
	}

	return true
}

// IsDisplayNameUsed returns true if the given display name is already in use
func (k Keeper) IsDisplayNameUsed(ctx sdk.Context, displayName string) bool {
	iter := k.StoreIterator(ctx, types.MarketKey, nil)
	defer iter.Close()
	var market types.Market
	for ; iter.Valid(); iter.Next() {
		k.Unmarshal(iter.Value(), &market)
		if market.DisplayName == displayName {
			return true
		}
	}
	return false
}

// SetCreatedMarket sets a newly created market, emits the new market event,
// and calls the appropriate callbacks.
func (k Keeper) SetCreatedMarket(ctx sdk.Context, market types.Market) {
	k.SetMarket(ctx, market)

	k.hooks.AfterMarketCreated(ctx, market)

	err := ctx.EventManager().EmitTypedEvents(&types.MarketEvent{
		Market: &market,
		Type:   string(transitiontypes.TransitionNew),
	})
	if err != nil {
		panic(err)
	}
}

// SetMarket set a specific market in the store
func (k Keeper) SetMarket(ctx sdk.Context, market types.Market) {
	store := k.Store(ctx, types.MarketKey)
	market.ExpiryTime = market.ExpiryTime.UTC() // stardardise ExpiryTime, used in validate unique market
	b := k.Marshal(&market)
	store.Set([]byte(market.Name), b)
}

// RemoveMarket removes a market from the store
func (k Keeper) RemoveMarket(ctx sdk.Context, name string) {
	store := k.Store(ctx, types.MarketKey)
	store.Delete([]byte(name))
}

// GetMarketNameSequence returns current sequence number + 1, and sets it in the store
func (k Keeper) GetMarketNameSequence(ctx sdk.Context) int64 {
	return int64(k.sek.GenerateSequenceNumber(ctx, types.MarketNameSequenceSuffixKey))
}
