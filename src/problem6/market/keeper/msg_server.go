package keeper

import (
	"context"

	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateMarket(goCtx context.Context, msg *types.MsgCreateMarket) (*types.MsgCreateMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	market, err := k.Keeper.CreateMarket(ctx, msg.MarketType, msg.Base, msg.Quote,
		msg.CurrentBasePriceUsd, msg.CurrentQuotePriceUsd, msg.IndexOracleId, msg.ExpiryTime)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateMarketResponse{
		Name: market.Name,
	}, nil
}
