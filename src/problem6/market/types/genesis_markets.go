package types

import (
	"time"

	"github.com/Switcheo/carbon/common/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var GenesisMarkets = []Market{
	{
		Name:                        "swth_eth",
		DisplayName:                 "SWTH_ETH",
		Description:                 "SWTH/ETH spot market",
		MarketType:                  "spot",
		Base:                        "swth",
		Quote:                       "eth",
		BasePrecision:               8,
		QuotePrecision:              18,
		LotSize:                     sdk.NewInt(10000000000),       // 100 SWTH
		TickSize:                    sdk.MustNewDecFromStr("1000"), // 0.0000001
		MinQuantity:                 sdk.NewInt(20000000000),       // 200 SWTH
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "swth_btc",
		DisplayName:                 "SWTH_BTC",
		Description:                 "SWTH/BTC spot market",
		MarketType:                  "spot",
		Base:                        "swth",
		Quote:                       "btc",
		BasePrecision:               8,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(10000000000), // 100 SWTH
		TickSize:                    sdk.MustNewDecFromStr("0.00000001"),
		MinQuantity:                 sdk.NewInt(20000000000), // 200 SWTH
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "btc_z29",
		DisplayName:                 "BTC_Z29",
		Description:                 "BTC/USD futures contract settling December 2029",
		MarketType:                  "futures",
		Base:                        "btc",
		Quote:                       "iusd",
		BasePrecision:               8,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(10000),             // 0.0001 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.25"), // $0.25
		MinQuantity:                 sdk.NewInt(10000),
		RiskStepSize:                sdk.NewInt(1000_0000),             // 0.1 btc
		InitialMarginBase:           sdk.MustNewDecFromStr("0.01"),     // 1% => 1 / 0.01 = 100x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.000005"), // increase by 0.005% every 0.1 btc
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.7"),      // liquidation begins when margin drops below 0.7% (at base risk)
		MaxLiquidationOrderTicket:   sdk.NewInt(111_0000_0000),         // 11 btc
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  sdk.NewInt(100000000), // 1 btc
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            20000, // +/- 100% (max for testing)
		IndexOracleId:               "DXBT",
		ExpiryTime:                  time.Date(2029, time.December, 31, 17, 0, 0, 0, time.UTC), // dec 31st 2029 5pm utc
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "eth_z29",
		DisplayName:                 "ETH_Z29",
		Description:                 "ETH/USD futures contract settling December 2029",
		MarketType:                  "futures",
		Base:                        "eth",
		Quote:                       "iusd",
		BasePrecision:               18,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(100000000000000),             // 0.01 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.000000000010"), // $0.10
		MinQuantity:                 sdk.NewInt(100000000000000),
		RiskStepSize:                sdk.NewInt(5000000000000000000),                     // 5 eth
		InitialMarginBase:           sdk.MustNewDecFromStr("0.02"),                       // 2% => 1 / 0.02 = 50x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.00001"),                    // increase by 0.001% every 5 eth
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.5"),                        // liquidation begins when margin drops below 1% (at base risk)
		MaxLiquidationOrderTicket:   utils.MustNewIntFromString("500000000000000000000"), // 500 eth
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  utils.MustNewIntFromString("50000000000000000000"), // 50 eth
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            20000, // +/- 100% (max for testing)
		IndexOracleId:               "DETH",
		ExpiryTime:                  time.Date(2029, time.December, 31, 17, 0, 0, 0, time.UTC), // dec 31st 2029 5pm utc
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "eth_z39",
		DisplayName:                 "ETH_Z39",
		Description:                 "ETH/USD futures contract settling December 2039",
		MarketType:                  "futures",
		Base:                        "eth",
		Quote:                       "iusd",
		BasePrecision:               18,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(100000000000000),             // 0.01 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.000000000010"), // $0.10
		MinQuantity:                 sdk.NewInt(100000000000000),
		RiskStepSize:                sdk.NewInt(5000000000000000000),                     // 5 eth
		InitialMarginBase:           sdk.MustNewDecFromStr("0.02"),                       // 2% => 1 / 0.02 = 50x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.00001"),                    // increase by 0.001% every 5 eth
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.5"),                        // liquidation begins when margin drops below 1% (at base risk)
		MaxLiquidationOrderTicket:   utils.MustNewIntFromString("500000000000000000000"), // 500 eth
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  utils.MustNewIntFromString("50000000000000000000"), // 50 eth
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            1500, // +/- 7.5%
		IndexOracleId:               "DETH",
		ExpiryTime:                  time.Date(2039, time.December, 31, 17, 0, 0, 0, time.UTC), // dec 31st 2029 5pm utc
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "eth_z19",
		DisplayName:                 "ETH_Z19",
		Description:                 "ETH/USD futures contract settling 15 mins past genesis",
		MarketType:                  "futures",
		Base:                        "eth",
		Quote:                       "iusd",
		BasePrecision:               18,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(100000000000000),             // 0.01 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.000000000010"), // $0.10
		MinQuantity:                 sdk.NewInt(100000000000000),
		RiskStepSize:                sdk.NewInt(5000000000000000000),                     // 5 eth
		InitialMarginBase:           sdk.MustNewDecFromStr("0.02"),                       // 2% => 1 / 0.02 = 50x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.00001"),                    // increase by 0.001% every 5 eth
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.5"),                        // liquidation begins when margin drops below 1% (at base risk)
		MaxLiquidationOrderTicket:   utils.MustNewIntFromString("500000000000000000000"), // 500 eth
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  utils.MustNewIntFromString("50000000000000000000"), // 50 eth
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            1500, // +/- 7.5%
		IndexOracleId:               "DETH",
		ExpiryTime:                  time.Now().Add(time.Minute * 15), // 15 minutes past genesis
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "eth_dai",
		DisplayName:                 "ETH_DAI",
		Description:                 "ETH/DAI spot market",
		MarketType:                  "spot",
		Base:                        "eth",
		Quote:                       "dai",
		BasePrecision:               18,
		QuotePrecision:              18,
		LotSize:                     sdk.NewInt(1000000000000000), // 0.001 ETH
		TickSize:                    sdk.MustNewDecFromStr("0.01"),
		MinQuantity:                 sdk.NewInt(1000000000000000), // 0.001 ETH
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "eth_usdc",
		DisplayName:                 "ETH_USDC",
		Description:                 "ETH/USDC spot market",
		MarketType:                  "spot",
		Base:                        "eth",
		Quote:                       "usdc",
		BasePrecision:               18,
		QuotePrecision:              6,
		LotSize:                     sdk.NewInt(1000000000000000),             // 0.001 ETH
		TickSize:                    sdk.MustNewDecFromStr("0.0000000000001"), // $0.10
		MinQuantity:                 sdk.NewInt(1000000000000000),             // 0.001 ETH
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "btc_z29u",
		DisplayName:                 "BTC_Z29U",
		Description:                 "BTC/USD futures contract USDC version",
		MarketType:                  "futures",
		Base:                        "btc",
		Quote:                       "usdc",
		BasePrecision:               8,
		QuotePrecision:              6,
		LotSize:                     sdk.NewInt(10000),               // 0.0001 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.0025"), // $0.25
		MinQuantity:                 sdk.NewInt(10000),
		RiskStepSize:                sdk.NewInt(1000_0000),             // 0.1 btc
		InitialMarginBase:           sdk.MustNewDecFromStr("0.01"),     // 1% => 1 / 0.01 = 100x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.000005"), // increase by 0.005% every 0.1 btc
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.7"),      // liquidation begins when margin drops below 0.7% (at base risk)
		MaxLiquidationOrderTicket:   sdk.NewInt(111_0000_0000),         // 11 btc
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  sdk.NewInt(100000000), // 1 btc
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            1500,
		IndexOracleId:               "DXBT",
		ExpiryTime:                  time.Now().Add(time.Hour * 24), // 24 hour past genesis
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "eth_z29u",
		DisplayName:                 "ETH_Z29U",
		Description:                 "ETH/USD futures contract USDC version",
		MarketType:                  "futures",
		Base:                        "eth",
		Quote:                       "usdc",
		BasePrecision:               18,
		QuotePrecision:              6,
		LotSize:                     sdk.NewInt(100000000000000),               // 0.01 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.00000000000010"), // $0.10
		MinQuantity:                 sdk.NewInt(100000000000000),
		RiskStepSize:                sdk.NewInt(5000000000000000000),                     // 5 eth
		InitialMarginBase:           sdk.MustNewDecFromStr("0.02"),                       // 2% => 1 / 0.02 = 50x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.00001"),                    // increase by 0.001% every 5 eth
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.5"),                        // liquidation begins when margin drops below 1% (at base risk)
		MaxLiquidationOrderTicket:   utils.MustNewIntFromString("500000000000000000000"), // 500 eth
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  utils.MustNewIntFromString("50000000000000000000"), // 50 eth
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            1500,
		IndexOracleId:               "DETH",
		ExpiryTime:                  time.Now().Add(time.Hour * 24), // 24 hour past genesis
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "btc_z19",
		DisplayName:                 "BTC_Z19",
		Description:                 "BTC/USD futures contract settling 15 mins past genesis",
		MarketType:                  "futures",
		Base:                        "btc",
		Quote:                       "iusd",
		BasePrecision:               8,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(10000),             // 0.0001 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.25"), // $0.25
		MinQuantity:                 sdk.NewInt(10000),
		RiskStepSize:                sdk.NewInt(1000_0000),             // 0.1 btc
		InitialMarginBase:           sdk.MustNewDecFromStr("0.01"),     // 1% => 1 / 0.01 = 100x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.000005"), // increase by 0.005% every 0.1 btc
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.7"),      // liquidation begins when margin drops below 0.7% (at base risk)
		MaxLiquidationOrderTicket:   sdk.NewInt(111_0000_0000),         // 11 btc
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  sdk.NewInt(100000000), // 1 btc
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            1500,
		IndexOracleId:               "DXBT",
		ExpiryTime:                  time.Now().Add(time.Hour * 24), // 24 hour past genesis
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "swth_nex",
		DisplayName:                 "SWTH_NEX",
		Description:                 "SWTH/NEX spot market",
		MarketType:                  "spot",
		Base:                        "swth",
		Quote:                       "nex",
		BasePrecision:               8,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(1000000), // 0.001 NEX
		TickSize:                    sdk.MustNewDecFromStr("0.001"),
		MinQuantity:                 sdk.NewInt(1000000), // 0.001 NEX
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "swth_nexo",
		DisplayName:                 "SWTH_NEXO",
		Description:                 "SWTH/NEXO spot market",
		MarketType:                  "spot",
		Base:                        "swth",
		Quote:                       "nexo",
		BasePrecision:               8,
		QuotePrecision:              18,
		LotSize:                     sdk.NewInt(1000000), // 0.01 NEXO
		TickSize:                    sdk.MustNewDecFromStr("1000000"),
		MinQuantity:                 sdk.NewInt(1000000), // 0.01 NEXO
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "swth_btc2",
		DisplayName:                 "SWTH_BTC2",
		Description:                 "SWTH/BTC2 spot market",
		MarketType:                  "spot",
		Base:                        "swth",
		Quote:                       "btc",
		BasePrecision:               8,
		QuotePrecision:              8,
		LotSize:                     sdk.NewInt(10000000000), // 100 SWTH
		TickSize:                    sdk.MustNewDecFromStr("0.00000001"),
		MinQuantity:                 sdk.NewInt(20000000000), // 200 SWTH
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "usdc_dai",
		DisplayName:                 "USDC_DAI",
		Description:                 "USDC/DAI spot market",
		MarketType:                  "spot",
		Base:                        "usdc",
		Quote:                       "dai",
		BasePrecision:               6,
		QuotePrecision:              18,
		LotSize:                     sdk.NewInt(1000000),                // 1 USDC
		TickSize:                    sdk.MustNewDecFromStr("100000000"), // price = quote/base (*12dp) target: 0.0001 * 10**12 = 10**8
		MinQuantity:                 sdk.NewInt(1000000),                // 1 USDC
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "BTC_PERP.USDC",
		DisplayName:                 "BTC_PERP.USDC",
		Description:                 "Bitcoin Futures Contract Perpetual",
		MarketType:                  "futures",
		Base:                        "btc",
		Quote:                       "usdc",
		BasePrecision:               8,
		QuotePrecision:              6,
		LotSize:                     sdk.NewInt(10000),               // 0.0001 quantity step
		TickSize:                    sdk.MustNewDecFromStr("0.0025"), // $0.25
		MinQuantity:                 sdk.NewInt(10000),
		RiskStepSize:                sdk.NewInt(1000_0000),             // 0.1 btc
		InitialMarginBase:           sdk.MustNewDecFromStr("0.01"),     // 1% => 1 / 0.01 = 100x
		InitialMarginStep:           sdk.MustNewDecFromStr("0.000005"), // increase by 0.005% every 0.1 btc
		MaintenanceMarginRatio:      sdk.MustNewDecFromStr("0.7"),      // liquidation begins when margin drops below 0.7% (at base risk)
		MaxLiquidationOrderTicket:   sdk.NewInt(111_0000_0000),         // 11 btc
		MaxLiquidationOrderDuration: 30 * time.Second,
		ImpactSize:                  sdk.NewInt(100000000), // 1 btc
		MarkPriceBand:               2000,
		LastPriceProtectedBand:      200,
		TradingBandwidth:            1500,
		IndexOracleId:               "DXBT",
		ExpiryTime:                  time.Unix(0, 0), // perps
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
	{
		Name:                        "swth_dai",
		DisplayName:                 "SWTH_DAI",
		Description:                 "SWTH/DAI spot market",
		MarketType:                  "spot",
		Base:                        "swth",
		Quote:                       "dai",
		BasePrecision:               8,
		QuotePrecision:              18,
		LotSize:                     sdk.NewInt(10000000000),     // 100 SWTH
		TickSize:                    sdk.MustNewDecFromStr("10"), // 0.00001
		MinQuantity:                 sdk.NewInt(20000000000),     // 200 SWTH
		RiskStepSize:                sdk.ZeroInt(),
		InitialMarginBase:           sdk.OneDec(), // 1x
		InitialMarginStep:           sdk.ZeroDec(),
		MaintenanceMarginRatio:      sdk.ZeroDec(),
		MaxLiquidationOrderTicket:   sdk.ZeroInt(),
		MaxLiquidationOrderDuration: 0,
		ImpactSize:                  sdk.ZeroInt(),
		MarkPriceBand:               0,
		LastPriceProtectedBand:      0,
		TradingBandwidth:            0,
		IndexOracleId:               "",
		ExpiryTime:                  time.Unix(0, 0),
		IsActive:                    true,
		IsSettled:                   false,
		CreatedBlockHeight:          0,
		ClosedBlockHeight:           0,
	},
}

var GenesisFeeStructures = []FeeStructure{
	// Market Type Fee Tiers
	{
		FeeCategory: FeeCategory{
			MarketType: "spot",
			MarketId:   "",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.003"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketType: "futures",
			MarketId:   "",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0.001"),
				// Changed for testing, need to be highest fee
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	// Market Ids Fee Tiers
	{
		FeeCategory: FeeCategory{
			MarketId: "swth_eth",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketId: "swth_btc",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketId: "btc_z29",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketId: "eth_z29",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.00025"),
				TakerFee: sdk.MustNewDecFromStr("0.00075"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketId: "eth_z39",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.0003"),
				TakerFee: sdk.MustNewDecFromStr("0.0003"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketId: "eth_z19",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
	{
		FeeCategory: FeeCategory{
			MarketId: "eth_dai",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "eth_usdc",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "btc_z29u",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "eth_z29u",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.00025"),
				TakerFee: sdk.MustNewDecFromStr("0.00075"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "btc_z19",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "swth_nex",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "swth_nexo",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "swth_btc2",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0.001"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "usdc_dai",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("0"),
				TakerFee: sdk.MustNewDecFromStr("0.002"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "BTC_PERP.USDC",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.0002"),
				TakerFee: sdk.MustNewDecFromStr("0.0005"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	}, {
		FeeCategory: FeeCategory{
			MarketId: "swth_dai",
		},
		FeeTiers: FeeTiers{{
			TradingFees: TradingFees{
				MakerFee: sdk.MustNewDecFromStr("-0.0002"),
				TakerFee: sdk.MustNewDecFromStr("0.0005"),
			},
			RequiredStake: sdk.ZeroInt(),
		}},
	},
}

var GenesisStakeEq = []StakeEquivalence{} // no stakeEquivalence set in case affect other tests
