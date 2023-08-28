package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	CreateMarketProposalType                    = "CreateMarket"
	UpdateMarketProposalType                    = "UpdateMarket"
	UpdatePerpetualsFundingIntervalProposalType = "UpdatePerpetualsFundingInterval"
)

func init() {
	govtypes.RegisterProposalType(CreateMarketProposalType) // retain for legacy
	govtypes.RegisterProposalType(UpdateMarketProposalType)
	govtypes.RegisterProposalType(UpdatePerpetualsFundingIntervalProposalType)

	govtypes.ModuleCdc.Amino.RegisterConcrete(&CreateMarketProposal{}, CreateMarketProposalAminoType, nil) // retain for legacy
	govtypes.ModuleCdc.Amino.RegisterConcrete(&UpdateMarketProposal{}, UpdateMarketProposalAminoType, nil)
	govtypes.ModuleCdc.Amino.RegisterConcrete(&UpdatePerpetualsFundingIntervalProposal{}, UpdatePerpetualsFundingIntervalProposalAminoType, nil)
}

var _ govtypes.Content = (*CreateMarketProposal)(nil)

func (p *CreateMarketProposal) GetTitle() string       { return p.Title }
func (p *CreateMarketProposal) GetDescription() string { return p.Description }
func (p *CreateMarketProposal) ProposalRoute() string  { return RouterKey }
func (p *CreateMarketProposal) ProposalType() string   { return CreateMarketProposalType }

func (p *CreateMarketProposal) ValidateBasic() error {
	switch "" {
	case p.Title:
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "title cannot be empty")
	case p.Description:
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be empty")
	}
	if err := p.Msg.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

var _ govtypes.Content = (*UpdateMarketProposal)(nil)

func NewUpdateMarketProposal(title, description string, msg MarketParams) *UpdateMarketProposal {
	return &UpdateMarketProposal{
		Title:       title,
		Description: description,
		Msg:         msg,
	}
}

func (p *UpdateMarketProposal) GetTitle() string       { return p.Title }
func (p *UpdateMarketProposal) GetDescription() string { return p.Description }
func (p *UpdateMarketProposal) ProposalRoute() string  { return RouterKey }
func (p *UpdateMarketProposal) ProposalType() string   { return UpdateMarketProposalType }

func (p *UpdateMarketProposal) ValidateBasic() error {
	switch "" {
	case p.Title:
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "title cannot be empty")
	case p.Description:
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be empty")
	}
	return nil
}

var _ govtypes.Content = (*UpdatePerpetualsFundingIntervalProposal)(nil)

func NewUpdatePerpetualsFundingIntervalProposal(title, description string, perpetualFundingInterval time.Duration) *UpdatePerpetualsFundingIntervalProposal {
	return &UpdatePerpetualsFundingIntervalProposal{
		Title:                     title,
		Description:               description,
		PerpetualsFundingInterval: perpetualFundingInterval,
	}
}

func (p *UpdatePerpetualsFundingIntervalProposal) GetTitle() string       { return p.Title }
func (p *UpdatePerpetualsFundingIntervalProposal) GetDescription() string { return p.Description }
func (p *UpdatePerpetualsFundingIntervalProposal) ProposalRoute() string  { return RouterKey }
func (p *UpdatePerpetualsFundingIntervalProposal) ProposalType() string {
	return UpdatePerpetualsFundingIntervalProposalType
}

func (p *UpdatePerpetualsFundingIntervalProposal) ValidateBasic() error {
	switch "" {
	case p.Title:
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "title cannot be empty")
	case p.Description:
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be empty")
	}

	if p.PerpetualsFundingInterval.Nanoseconds()%time.Hour.Nanoseconds() != 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "perpetuals funding interval must be in hours resolution")
	}

	if p.PerpetualsFundingInterval < time.Hour {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "perpetuals funding interval must be at least 1 hour")
	}

	return nil
}
