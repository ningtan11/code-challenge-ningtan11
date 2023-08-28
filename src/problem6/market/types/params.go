package types

import (
	fmt "fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	yaml "gopkg.in/yaml.v2"
)

var DefaultLotSizeUsd = sdk.MustNewDecFromStr("1")
var DefaultTickSizeUsd = sdk.MustNewDecFromStr("0.001") // relative to a base price of $1 (i.e this is the minimum spread %)
var DefaultMinQuantityUsd = sdk.MustNewDecFromStr("1")
var DefaultRiskStepSizeUsd = sdk.MustNewDecFromStr("100000")
var DefaultInitialMarginBase = sdk.MustNewDecFromStr("0.2")
var DefaultInitialMarginStep = sdk.MustNewDecFromStr("0.0001")
var DefaultMaxLiquidationOrderTicketUsd = sdk.MustNewDecFromStr("100000")
var DefaultSpotMakerFee = sdk.MustNewDecFromStr("-0.001")
var DefaultFuturesMakerFee = sdk.MustNewDecFromStr("-0.0002")
var DefaultSpotTakerFee = sdk.MustNewDecFromStr("0.002")
var DefaultFuturesTakerFee = sdk.MustNewDecFromStr("0.0005")
var DefaultMaintenanceMarginRatio = sdk.MustNewDecFromStr("0.7")
var DefaultMaxLiquidationOrderDuration = 60 * time.Second
var DefaultImpactSizeUsd = sdk.MustNewDecFromStr("100000") // approx 1 btc
var DefaultMarkPriceBand = uint32(100)
var DefaultLastPriceProtectedBand = uint32(200)
var MaxActiveMarkets = uint32(300)
var DefaultTradingBandwidth = uint32(300)

var DefaultFundingRateBand = sdk.MustNewDecFromStr("175.20") // 2% hourly = 2*24*365 = 17520% yearly = 175.20x
var FundingRateBandDurationSeconds = sdk.NewDec(31536000)    // 1 year in seconds

var DefaultLpSpotMakerFee = sdk.ZeroDec()
var DefaultLpSpotTakerFee = sdk.ZeroDec()
var DefaultLpFuturesMakerFee = sdk.ZeroDec()
var DefaultLpFuturesTakerFee = sdk.ZeroDec()

var (
	KeyLotSize                     = []byte("DefaultLotSizeUsd")
	KeyTickSize                    = []byte("DefaultTickSizeUsd")
	KeyMinQuantity                 = []byte("DefaultMinQuantityUsd")
	KeySpotMakerFee                = []byte("DefaultSpotMakerFee")
	KeyFuturesMakerFee             = []byte("DefaultFuturesMakerFee")
	KeySpotTakerFee                = []byte("DefaultSpotTakerFee")
	KeyFuturesTakerFee             = []byte("DefaultFuturesTakerFee")
	KeyRiskStepSize                = []byte("DefaultRiskStepSizeUsd")
	KeyInitialMarginBase           = []byte("DefaultInitialMarginBase")
	KeyInitialMarginStep           = []byte("DefaultInitialMarginStep")
	KeyMaintenanceMarginRatio      = []byte("DefaultMaintenanceMarginRatio")
	KeyMaxLiquidationOrderTicket   = []byte("DefaultMaxLiquidationOrderTicketUsd")
	KeyMaxLiquidationOrderDuration = []byte("DefaultMaxLiquidationOrderDuration")
	KeyImpactSize                  = []byte("DefaultImpactSize")
	KeyMarkPriceBand               = []byte("DefaultMarkPriceBand")
	KeyLastPriceProtectedBand      = []byte("DefaultLastPriceProtectedBand")
	KeyMaxActiveMarkets            = []byte("MaxActiveMarkets")
	KeyTradingBandwidth            = []byte("TradingBandwidth")
	KeyFundingRateBand             = []byte("FundingRateBand") // in units of 365days
	KeyLpSpotMakerFee              = []byte("DefaultLPSpotMakerFee")
	KeyLpSpotTakerFee              = []byte("DefaultLPSpotTakerFee")
	KeyLpFuturesMakerFee           = []byte("DefaultLPFuturesMakerFee")
	KeyLpFuturesTakerFee           = []byte("DefaultLPFuturesTakerFee")
)

// ParamKeyTable for liquiditypool module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(defaultLotSizeUsd sdk.Dec, defaultTickSizeUsd sdk.Dec,
	defaultMinQuantityUsd sdk.Dec, defaultRiskStepSizeUsd sdk.Dec,
	defaultInitialMarginBase sdk.Dec, defaultInitialMarginStep sdk.Dec,
	defaultMaxLiquidationOrderTicketUsd sdk.Dec, defaultSpotMakerFee sdk.Dec,
	defaultFuturesMakerFee sdk.Dec, defaultSpotTakerFee sdk.Dec,
	defaultFuturesTakerFee sdk.Dec, defaultMaintenanceMarginRatio sdk.Dec,
	defaultMaxLiquidationOrderDuration time.Duration, defaultImpactSizeUsd sdk.Dec,
	defaultMarkPriceBand uint32, defaultLastPriceProtectedBand uint32, maxActiveMarkets uint32,
	defaultTradingBandwidth uint32, fundingRateBand sdk.Dec, defaultLpSpotMakerFee sdk.Dec, defaultLpSpotTakerFee sdk.Dec, defaultLpFuturesMakerFee sdk.Dec, defaultLpFuturesTakerFee sdk.Dec) Params {
	return Params{
		DefaultLotSizeUsd:                   defaultLotSizeUsd,
		DefaultTickSizeUsd:                  defaultTickSizeUsd,
		DefaultMinQuantityUsd:               defaultMinQuantityUsd,
		DefaultRiskStepSizeUsd:              defaultRiskStepSizeUsd,
		DefaultInitialMarginBase:            defaultInitialMarginBase,
		DefaultInitialMarginStep:            defaultInitialMarginStep,
		DefaultMaxLiquidationOrderTicketUsd: defaultMaxLiquidationOrderTicketUsd,
		DefaultFuturesMakerFee:              defaultFuturesMakerFee,
		DefaultSpotMakerFee:                 defaultSpotMakerFee,
		DefaultSpotTakerFee:                 defaultSpotTakerFee,
		DefaultFuturesTakerFee:              defaultFuturesTakerFee,
		DefaultMaintenanceMarginRatio:       defaultMaintenanceMarginRatio,
		DefaultMaxLiquidationOrderDuration:  defaultMaxLiquidationOrderDuration,
		DefaultImpactSizeUsd:                defaultImpactSizeUsd,
		DefaultMarkPriceBand:                defaultMarkPriceBand,
		DefaultLastPriceProtectedBand:       defaultLastPriceProtectedBand,
		MaxActiveMarkets:                    maxActiveMarkets,
		DefaultTradingBandwidth:             defaultTradingBandwidth,
		FundingRateBand:                     fundingRateBand,
		DefaultLpSpotMakerFee:               defaultLpSpotMakerFee,
		DefaultLpSpotTakerFee:               defaultLpSpotTakerFee,
		DefaultLpFuturesMakerFee:            defaultLpFuturesMakerFee,
		DefaultLpFuturesTakerFee:            defaultLpFuturesTakerFee,
	}
}

func DefaultParams() Params {
	return Params{
		DefaultLotSizeUsd:                   DefaultLotSizeUsd,
		DefaultTickSizeUsd:                  DefaultTickSizeUsd,
		DefaultMinQuantityUsd:               DefaultMinQuantityUsd,
		DefaultRiskStepSizeUsd:              DefaultRiskStepSizeUsd,
		DefaultInitialMarginBase:            DefaultInitialMarginBase,
		DefaultInitialMarginStep:            DefaultInitialMarginStep,
		DefaultMaxLiquidationOrderTicketUsd: DefaultMaxLiquidationOrderTicketUsd,
		DefaultFuturesMakerFee:              DefaultFuturesMakerFee,
		DefaultSpotMakerFee:                 DefaultSpotMakerFee,
		DefaultSpotTakerFee:                 DefaultSpotTakerFee,
		DefaultFuturesTakerFee:              DefaultFuturesTakerFee,
		DefaultMaintenanceMarginRatio:       DefaultMaintenanceMarginRatio,
		DefaultMaxLiquidationOrderDuration:  DefaultMaxLiquidationOrderDuration,
		DefaultImpactSizeUsd:                DefaultImpactSizeUsd,
		DefaultMarkPriceBand:                DefaultMarkPriceBand,
		DefaultLastPriceProtectedBand:       DefaultLastPriceProtectedBand,
		MaxActiveMarkets:                    MaxActiveMarkets,
		DefaultTradingBandwidth:             DefaultTradingBandwidth,
		FundingRateBand:                     DefaultFundingRateBand,
		DefaultLpSpotMakerFee:               DefaultLpSpotMakerFee,
		DefaultLpSpotTakerFee:               DefaultLpSpotTakerFee,
		DefaultLpFuturesMakerFee:            DefaultLpFuturesMakerFee,
		DefaultLpFuturesTakerFee:            DefaultLpFuturesTakerFee,
	}
}

// Validate all bank module parameters
func (p Params) Validate() error {
	return nil
	// return validateLotSize(p.LotSize)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyLotSize, &p.DefaultLotSizeUsd, validateLotSize),
		paramstypes.NewParamSetPair(KeyTickSize, &p.DefaultTickSizeUsd, validateTickSize),
		paramstypes.NewParamSetPair(KeyMinQuantity, &p.DefaultMinQuantityUsd, validateMinQuantity),
		paramstypes.NewParamSetPair(KeySpotMakerFee, &p.DefaultSpotMakerFee, validateFee),
		paramstypes.NewParamSetPair(KeyFuturesMakerFee, &p.DefaultFuturesMakerFee, validateFee),
		paramstypes.NewParamSetPair(KeySpotTakerFee, &p.DefaultSpotTakerFee, validateFee),
		paramstypes.NewParamSetPair(KeyFuturesTakerFee, &p.DefaultFuturesTakerFee, validateFee),
		paramstypes.NewParamSetPair(KeyRiskStepSize, &p.DefaultRiskStepSizeUsd, validateRiskStepSize),
		paramstypes.NewParamSetPair(KeyInitialMarginBase, &p.DefaultInitialMarginBase, validateInitialMarginBase),
		paramstypes.NewParamSetPair(KeyInitialMarginStep, &p.DefaultInitialMarginStep, validateInitialMarginStep),
		paramstypes.NewParamSetPair(KeyMaintenanceMarginRatio, &p.DefaultMaintenanceMarginRatio, validateMaintenanceMarginRatio),
		paramstypes.NewParamSetPair(KeyMaxLiquidationOrderTicket, &p.DefaultMaxLiquidationOrderTicketUsd, validateMaxLiquidationOrderTicket),
		paramstypes.NewParamSetPair(KeyMaxLiquidationOrderDuration, &p.DefaultMaxLiquidationOrderDuration, validateMaxLiquidationOrderDuration),
		paramstypes.NewParamSetPair(KeyImpactSize, &p.DefaultImpactSizeUsd, validateImpactSize),
		paramstypes.NewParamSetPair(KeyMarkPriceBand, &p.DefaultMarkPriceBand, validateMarkPriceBand),
		paramstypes.NewParamSetPair(KeyLastPriceProtectedBand, &p.DefaultLastPriceProtectedBand, validateLastPriceProtectedBand),
		paramstypes.NewParamSetPair(KeyMaxActiveMarkets, &p.MaxActiveMarkets, validateMaxActiveMarkets),
		paramstypes.NewParamSetPair(KeyTradingBandwidth, &p.DefaultTradingBandwidth, validateTradingBandwidth),
		paramstypes.NewParamSetPair(KeyFundingRateBand, &p.FundingRateBand, validateFundingRateBand),
		paramstypes.NewParamSetPair(KeyLpSpotMakerFee, &p.DefaultLpSpotMakerFee, validateLPMakerFee),
		paramstypes.NewParamSetPair(KeyLpSpotTakerFee, &p.DefaultLpSpotTakerFee, validateLPTakerFee),
		paramstypes.NewParamSetPair(KeyLpFuturesMakerFee, &p.DefaultLpFuturesMakerFee, validateLPMakerFee),
		paramstypes.NewParamSetPair(KeyLpFuturesTakerFee, &p.DefaultLpFuturesTakerFee, validateLPTakerFee),
	}
}

func validateLotSize(i interface{}) error {
	lotSize, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	switch {
	case !lotSize.IsPositive():
		return fmt.Errorf("lot_size must be more than zero: %v", i)
	case lotSize.GT(sdk.NewDecFromInt(MaxLotSize)):
		return fmt.Errorf("lot_size is too large: %v", i)
	}
	return nil
}

func validateTickSize(i interface{}) error {
	tickSize, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	switch {
	case !tickSize.IsPositive():
		return fmt.Errorf("tick_size must be positive, %v", i)
	case tickSize.GT(MaxTickSize):
		return fmt.Errorf("tick_size is too large, %v", i)
	}
	return nil
}

func validateMinQuantity(i interface{}) error {
	minQuantity, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	switch {
	case !minQuantity.IsPositive():
		return fmt.Errorf("min_quantity must be positive, %v", i)
	}
	return nil
}

func validateFee(i interface{}) error {
	fee, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	switch {
	case fee.LTE(sdk.OneDec().MulInt64(-1)):
		return fmt.Errorf("fee cannot be less than or equal to -1, %v", i)
	case fee.GTE(sdk.OneDec()):
		return fmt.Errorf("fee cannot be greater or equal to 1, %v", i)
	case !fee.Quo(sdk.MustNewDecFromStr("0.00001")).IsInteger():
		return fmt.Errorf("taker fee can only have up to 0.1 bps precision, %v", i)
	}
	return nil
}

func validateRiskStepSize(i interface{}) error {
	riskStepSize, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if riskStepSize.IsNegative() {
		return fmt.Errorf("risk_step_size must not be negative, %v", i)
	}
	return nil
}

func validateInitialMarginBase(i interface{}) error {
	initialMarginBase, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if initialMarginBase.IsNegative() {
		return fmt.Errorf("initial_margin_base must not be negative, %v", i)
	}
	return nil
}

func validateInitialMarginStep(i interface{}) error {
	initialMarginStep, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if initialMarginStep.IsNegative() {
		return fmt.Errorf("initial_margin_step must not be negative, %v", i)
	}
	return nil
}

func validateMaintenanceMarginRatio(i interface{}) error {
	maintenanceMarginRatio, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if !maintenanceMarginRatio.IsPositive() || maintenanceMarginRatio.GTE(sdk.OneDec()) {
		return fmt.Errorf("maintenance_margin_ratio must between zero and one, %v", i)
	}
	return nil
}

func validateMaxLiquidationOrderTicket(i interface{}) error {
	_, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	return nil
}

func validateMaxLiquidationOrderDuration(i interface{}) error {
	maxLiquidationOrderDuration, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if maxLiquidationOrderDuration.Seconds() < 3 {
		return fmt.Errorf("max_liquidation_order_duration must be at least 30 seconds, %v", i)
	}
	return nil
}

func validateImpactSize(i interface{}) error {
	return nil
}

func validateMarkPriceBand(i interface{}) error {
	markPriceBand, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if markPriceBand == 0 || markPriceBand > 20000 {
		return fmt.Errorf("mark_price_band must be between 0 and 20000, %v", i)
	}
	return nil
}

func validateLastPriceProtectedBand(i interface{}) error {
	lastPriceProtectedBand, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if lastPriceProtectedBand == 0 || lastPriceProtectedBand > 20000 {
		return fmt.Errorf("last_price_protected_band must be between 0 and 20000, %v", i)
	}
	return nil
}

func validateMaxActiveMarkets(i interface{}) error {
	maxActiveMarkets, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if maxActiveMarkets == 0 {
		return fmt.Errorf("active_markets must be greater than 0, %v", i)
	}
	return nil
}

func validateTradingBandwidth(i interface{}) error {
	tradingBandwidth, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if tradingBandwidth == 0 {
		return fmt.Errorf("trading_bandwidth must be greater than 0, %v", i)
	}
	return nil
}

func validateFundingRateBand(i interface{}) error {
	fundingRateBand, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %v", i)
	}
	if fundingRateBand.IsNegative() {
		return fmt.Errorf("initial_margin_step must not be negative, %v", i)
	}
	return nil
}

func validateLPMakerFee(i interface{}) error {
	fee, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if fee.IsPositive() {
		return fmt.Errorf("maker fee should not be positive: %T", i)
	}
	return nil
}
func validateLPTakerFee(i interface{}) error {
	fee, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !fee.IsZero() {
		return fmt.Errorf("taker fee should only be zero: %T", i)
	}

	return nil
}
