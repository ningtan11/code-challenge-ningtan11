package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	CreateMarketProposalAminoType                    = "market/CreateMarketProposal"
	UpdateMarketProposalAminoType                    = "market/UpdateMarketProposal"
	UpdatePerpetualsFundingIntervalProposalAminoType = "market/UpdatePerpetualsFundingIntervalProposal"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDisableSpotMarket{}, "market/DisableSpotMarket", nil)
	cdc.RegisterConcrete(&MsgAddFeeTier{}, "market/AddFeeTier", nil)
	cdc.RegisterConcrete(&MsgRemoveFeeTier{}, "market/RemoveFeeTier", nil)
	cdc.RegisterConcrete(&MsgUpdateFeeTier{}, "market/UpdateFeeTier", nil)
	cdc.RegisterConcrete(&MsgSetStakeEquivalence{}, "market/SetStakeEquivalence", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateMarket{}, "market/CreateMarket", nil)
	cdc.RegisterConcrete(&MsgUpdateMarket{}, "market/UpdateMarket", nil)
	cdc.RegisterConcrete(&MsgUpdatePerpetualsFundingInterval{}, "market/UpdatePerpetualsFundingInterval", nil)
	// retain for legacy otherwise iteration for /cosmos/gov/v1beta1/proposals will fail
	cdc.RegisterConcrete(&CreateMarketProposal{}, CreateMarketProposalAminoType, nil)
	cdc.RegisterConcrete(&UpdateMarketProposal{}, UpdateMarketProposalAminoType, nil)
	cdc.RegisterConcrete(&UpdatePerpetualsFundingIntervalProposal{}, UpdatePerpetualsFundingIntervalProposalAminoType, nil)
	cdc.RegisterConcrete(&MsgUpdateAllPoolTradingFees{}, "market/UpdateAllPoolTradingFees", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMarket{},
		&MsgUpdateMarket{},
		&MsgUpdatePerpetualsFundingInterval{},
		&MsgDisableSpotMarket{},
		&MsgAddFeeTier{},
		&MsgRemoveFeeTier{},
		&MsgUpdateFeeTier{},
		&MsgSetStakeEquivalence{},
		&MsgUpdateAllPoolTradingFees{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&CreateMarketProposal{}, // retain for legacy
		&UpdateMarketProposal{},
		&UpdatePerpetualsFundingIntervalProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/market module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/market and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
