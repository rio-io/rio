package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateResume = "create_resume"

var _ sdk.Msg = &MsgCreateResume{}

func NewMsgCreateResume(creator string, certs []uint64, avatarUrl string, name string, description string) *MsgCreateResume {
	return &MsgCreateResume{
		Creator:     creator,
		Certs:       certs,
		AvatarUrl:   avatarUrl,
		Name:        name,
		Description: description,
	}
}

func (msg *MsgCreateResume) Route() string {
	return RouterKey
}

func (msg *MsgCreateResume) Type() string {
	return TypeMsgCreateResume
}

func (msg *MsgCreateResume) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateResume) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateResume) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
