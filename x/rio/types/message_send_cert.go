package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendCert = "send_cert"

var _ sdk.Msg = &MsgSendCert{}

func NewMsgSendCert(creator string, to string, certType string, title string, description string) *MsgSendCert {
	return &MsgSendCert{
		Creator:     creator,
		To:          to,
		CertType:    certType,
		Title:       title,
		Description: description,
	}
}

func (msg *MsgSendCert) Route() string {
	return RouterKey
}

func (msg *MsgSendCert) Type() string {
	return TypeMsgSendCert
}

func (msg *MsgSendCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
