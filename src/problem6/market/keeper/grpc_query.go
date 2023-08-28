package keeper

import (
	"context"
	"net/url"

	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Market(c context.Context, req *types.QueryGetMarketRequest) (*types.QueryGetMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	reqName, err := url.QueryUnescape(req.Name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	market, found := k.GetMarket(ctx, reqName)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMarketResponse{Market: market}, nil
}

/*
Example output from the grpc query above:
{
  "market":
    {
      "name": "BTC_PERP.USDC",
      "display_name": "BTC_PERP.USDC",
      "description": "Bitcoin Futures Contract Perpetual",
      "market_type": "futures",
      "base": "btc",
      "quote": "usdc",
      "base_precision": "8",
      "quote_precision": "6",
      "lot_size": "10000",
      "tick_size": "0.002500000000000000",
      "min_quantity": "10000",
      "created_block_height": "0",
      "risk_step_size": "10000000",
      "initial_margin_base": "0.010000000000000000",
      "initial_margin_step": "0.000005000000000000",
      "maintenance_margin_ratio": "0.700000000000000000",
      "max_liquidation_order_ticket": "11100000000",
      "max_liquidation_order_duration": "30s",
      "impact_size": "100000000",
      "mark_price_band": 2000,
      "last_price_protected_band": 200,
      "index_oracle_id": "DXBT",
      "expiry_time": "1970-01-01T00:00:00Z",
      "is_active": true,
      "is_settled": false,
      "closed_block_height": "0",
      "trading_bandwidth": 1500
    }
	}
*/

func (k Keeper) MarketAll(c context.Context, req *types.QueryAllMarketRequest) (*types.QueryAllMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	markets, err := k.GetAllMarkets(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMarketResponse{Markets: markets}, nil
}
