package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateCert = "create_cert"

var _ sdk.Msg = &MsgCreateCert{}

func NewMsgCreateCert(creator string, title string) *MsgCreateCert {
	return &MsgCreateCert{
		Creator: creator,
		Title:   title,
	}
}

func (msg *MsgCreateCert) Route() string {
	return RouterKey
}

func (msg *MsgCreateCert) Type() string {
	return TypeMsgCreateCert
}

func (msg *MsgCreateCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
