package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type PricingKeeper interface {
	ResetLastFundingAt(ctx sdk.Context, key string, value time.Time)
	SetLastFundingAt(ctx sdk.Context, key string, value time.Time)
}

// ParamSubspace defines the expected Subspace interfacace
type ParamSubspace interface {
	WithKeyTable(table paramstypes.KeyTable) paramstypes.Subspace
	Get(ctx sdk.Context, key []byte, ptr interface{})
	GetParamSet(ctx sdk.Context, ps paramstypes.ParamSet)
	SetParamSet(ctx sdk.Context, ps paramstypes.ParamSet)
	GetParamSetIfExists(ctx sdk.Context, ps paramstypes.ParamSet)
	HasKeyTable() bool
}

type LiquidationKeeper interface {
	HasOutstandingPositionsForMarket(ctx sdk.Context, market string) bool
}

type LiquidityPoolKeeper interface {
	GetPoolIDsWithDenoms(ctx sdk.Context, denomOne, denomTwo string) (pools []uint64)
	AddPoolRoute(ctx sdk.Context, market string, poolIds []uint64, numQuotes int64) error
	PoolRouteExists(ctx sdk.Context, market string, poolIDs []uint64) bool
	GetParamNumQuotes(ctx sdk.Context) uint64
	GetAllPoolIDs(ctx sdk.Context) (pools []uint64)
}

type SequenceKeeper interface {
	GenerateSequenceNumber(ctx sdk.Context, suffix []byte) uint64
	SetSequenceNumber(ctx sdk.Context, key []byte, sequenceNumber uint64)
}

type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address
	GetDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) (delegations []stakingtypes.Delegation)
}

type AdminKeeper interface {
	IsAdmin(ctx sdk.Context, account string) bool
}
