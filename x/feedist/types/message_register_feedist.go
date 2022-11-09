package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterFeedist = "register_feedist"

var _ sdk.Msg = &MsgRegisterFeedist{}

func NewMsgRegisterFeedist(creator string, contract string, shares string) *MsgRegisterFeedist {
	return &MsgRegisterFeedist{
		Creator:  creator,
		Contract: contract,
		Shares:   shares,
	}
}

func (msg *MsgRegisterFeedist) Route() string {
	return RouterKey
}

func (msg *MsgRegisterFeedist) Type() string {
	return TypeMsgRegisterFeedist
}

func (msg *MsgRegisterFeedist) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterFeedist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterFeedist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
