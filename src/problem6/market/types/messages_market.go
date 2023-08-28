package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMarket{}

func NewMsgCreateMarket(creator string, marketType string, base string, quote string,
	currentBasePriceUsd sdk.Dec, currentQuotePriceUsd sdk.Dec, indexOracleId string, expiryTime time.Time) *MsgCreateMarket {
	return &MsgCreateMarket{
		Creator:              creator,
		MarketType:           marketType,
		Base:                 base,
		Quote:                quote,
		CurrentBasePriceUsd:  currentBasePriceUsd,
		CurrentQuotePriceUsd: currentQuotePriceUsd,
		IndexOracleId:        indexOracleId,
		ExpiryTime:           expiryTime,
	}
}

// Route ...
func (msg MsgCreateMarket) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgCreateMarket) Type() string {
	return "create_market"
}

// GetSigners ...
func (msg *MsgCreateMarket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *MsgCreateMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *MsgCreateMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "creator: %v", err)
	}
	switch {
	case msg.Base == "":
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Base token must be present")
	case msg.Quote == "":
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Base token must be present")
	case !msg.CurrentBasePriceUsd.IsPositive():
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Current base price must be positive")
	case !msg.CurrentQuotePriceUsd.IsPositive():
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Current quote price must be positive")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMarket{}

func NewMsgUpdateMarket(updater string, params MarketParams) *MsgUpdateMarket {
	return &MsgUpdateMarket{
		Updater:      updater,
		MarketParams: params,
	}
}

// Route ...
func (msg MsgUpdateMarket) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgUpdateMarket) Type() string {
	return "update_market"
}

// GetSigners ...
func (msg *MsgUpdateMarket) GetSigners() []sdk.AccAddress {
	updater, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{updater}
}

// GetSignBytes ...
func (msg *MsgUpdateMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *MsgUpdateMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "creator: %v", err)
	}
	return nil
}

// MsgUpdateInterval

var _ sdk.Msg = &MsgUpdatePerpetualsFundingInterval{}

func NewMsgUpdatePerpetualsFundingInterval(updater string, perpetuals_funding_interval time.Duration) *MsgUpdatePerpetualsFundingInterval {
	return &MsgUpdatePerpetualsFundingInterval{
		Updater:                   updater,
		PerpetualsFundingInterval: perpetuals_funding_interval,
	}
}

// Route ...
func (msg MsgUpdatePerpetualsFundingInterval) Route() string {
	return RouterKey
}

// Type ...
func (msg MsgUpdatePerpetualsFundingInterval) Type() string {
	return "update_perpetuals_funding_interval"
}

// GetSigners ...
func (msg *MsgUpdatePerpetualsFundingInterval) GetSigners() []sdk.AccAddress {
	updater, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{updater}
}

// GetSignBytes ...
func (msg *MsgUpdatePerpetualsFundingInterval) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *MsgUpdatePerpetualsFundingInterval) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "creator: %v", err)
	}
	if msg.PerpetualsFundingInterval.Nanoseconds()%time.Hour.Nanoseconds() != 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "perpetuals funding interval must be in hours resolution")
	}
	if msg.PerpetualsFundingInterval < time.Hour {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "perpetuals funding interval must be at least 1 hour")
	}
	return nil
}
