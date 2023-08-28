package market

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Switcheo/carbon/x/market/keeper"
	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func NewMarketProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, c govtypes.Content) (err error) {
		switch p := c.(type) {
		case *types.UpdateMarketProposal:
			return k.UpdateMarket(ctx, p.Msg)
		case *types.UpdatePerpetualsFundingIntervalProposal:
			k.UpdatePerpetualsFundingInterval(ctx, p.PerpetualsFundingInterval)
			return nil
		default:
			err = errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized market proposal content type: %T", p)
		}
		return
	}
}
