package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteName{}

// MsgDeleteName is a struct
type MsgDeleteName struct {
	Name  string         `json:"name" yaml:"name"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
}

// NewMsgDeleteName does stuff
func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}

// Route implements Msg interface
func (msg MsgDeleteName) Route() string {
	return RouterKey
}

// Type implements Msg interface
func (msg MsgDeleteName) Type() string {
	return "DeleteName"
}

// GetSigners implements Msg interface
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// GetSignBytes implements Msg interface
func (msg MsgDeleteName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg interface
func (msg MsgDeleteName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
