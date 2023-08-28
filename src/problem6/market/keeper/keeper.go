package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/Switcheo/carbon/x/market/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # ibc/keeper/import
	coinkeeper "github.com/Switcheo/carbon/x/coin/keeper"
	oraclekeeper "github.com/Switcheo/carbon/x/oracle/keeper"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

type (
	Keeper struct {
		cdc       codec.BinaryCodec
		storeKey  storetypes.StoreKey
		tStoreKey storetypes.StoreKey
		memKey    storetypes.StoreKey
		hooks     *types.Hooks
		// this line is used by starport scaffolding # ibc/keeper/attribute
		ck         coinkeeper.Keeper
		ok         oraclekeeper.Keeper
		prk        types.PricingKeeper
		lqk        types.LiquidationKeeper
		lpk        types.LiquidityPoolKeeper
		sek        types.SequenceKeeper
		stk        types.StakingKeeper
		adk        types.AdminKeeper
		paramspace types.ParamSubspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	tStoreKey,
	memKey storetypes.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	ck coinkeeper.Keeper,
	ok oraclekeeper.Keeper,
	prk types.PricingKeeper,
	lqk types.LiquidationKeeper,
	lpk types.LiquidityPoolKeeper,
	sek types.SequenceKeeper,
	stk types.StakingKeeper,
	adk types.AdminKeeper,
	ps types.ParamSubspace,
) *Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	return &Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		tStoreKey: tStoreKey,
		memKey:    memKey,
		hooks:     &types.Hooks{},
		// this line is used by starport scaffolding # ibc/keeper/return
		ck:         ck,
		ok:         ok,
		prk:        prk,
		lqk:        lqk,
		lpk:        lpk,
		sek:        sek,
		stk:        stk,
		adk:        adk,
		paramspace: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Store fetches the permanent store
func (k Keeper) Store(ctx sdk.Context, key string) prefix.Store {
	mainStore := ctx.KVStore(k.storeKey)
	return prefix.NewStore(mainStore, types.KeyPrefix(key))
}

// TStore fetches the transient store
func (k Keeper) TStore(ctx sdk.Context, key string) prefix.Store {
	mainStore := ctx.TransientStore(k.tStoreKey)
	return prefix.NewStore(mainStore, types.KeyPrefix(key))
}

// StoreIterator returns the iterator for the store
func (k Keeper) StoreIterator(ctx sdk.Context, key string, prefix []byte) sdk.Iterator {
	store := k.Store(ctx, key)
	return sdk.KVStorePrefixIterator(store, prefix)
}

// ReverseStoreIterator returns the reverse iterator for the store
func (k Keeper) ReverseStoreIterator(ctx sdk.Context, key string, prefix []byte) sdk.Iterator {
	store := k.Store(ctx, key)
	return sdk.KVStoreReversePrefixIterator(store, prefix)
}

// Marshal - marshals data into []byte, returns []byte
func (k Keeper) Marshal(val codec.ProtoMarshaler) []byte { return k.cdc.MustMarshal(val) }

// Unmarshal - unmarshals data into struct via ptr
func (k Keeper) Unmarshal(bytes []byte, ptr codec.ProtoMarshaler) {
	k.cdc.MustUnmarshal(bytes, ptr)
}
