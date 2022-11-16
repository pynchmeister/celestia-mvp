package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUploadW3S = "upload_w_3_s"

var _ sdk.Msg = &MsgUploadW3S{}

func NewMsgUploadW3S(creator string, title string, content string) *MsgUploadW3S {
	return &MsgUploadW3S{
		Creator: creator,
		Title:   title,
		Content: content,
	}
}

func (msg *MsgUploadW3S) Route() string {
	return RouterKey
}

func (msg *MsgUploadW3S) Type() string {
	return TypeMsgUploadW3S
}

func (msg *MsgUploadW3S) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUploadW3S) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUploadW3S) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
