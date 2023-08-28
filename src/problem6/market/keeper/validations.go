package keeper

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateMarket - shared validate market to share validate logic for create/update market
func (k Keeper) ValidateMarket(ctx sdk.Context, market types.Market) error {
	if market.MarketType == types.FuturesMarket {
		if !k.ok.HasOracle(ctx, market.IndexOracleId) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index oracle does not exist")
		}
		if !market.IsSettled && market.ExpiryTime.Unix() != 0 && market.ExpiryTime.Before(ctx.BlockTime().Add(time.Hour)) {
			return errorsmod.Wrap(types.ErrInvalidExpiry, market.ExpiryTime.String())
		}
	}

	return nil
}
