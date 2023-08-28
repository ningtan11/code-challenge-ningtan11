package keeper

import (
	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return
}

func (k Keeper) GetDefaultPoolTradingFees(ctx sdk.Context, marketType string) (fees types.TradingFees) {
	var makerFee, takerFee sdk.Dec
	switch marketType {
	case "spot":
		makerFee = k.GetParams(ctx).DefaultLpSpotMakerFee
		takerFee = k.GetParams(ctx).DefaultLpSpotTakerFee
	case "futures":
		makerFee = k.GetParams(ctx).DefaultLpFuturesMakerFee
		takerFee = k.GetParams(ctx).DefaultLpFuturesTakerFee
	}

	return types.TradingFees{MakerFee: makerFee, TakerFee: takerFee}
}
