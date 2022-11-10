package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethermint "github.com/evmos/ethermint/types"
)

const TypeMsgRegisterFeedist = "register"

var _ sdk.Msg = &MsgRegisterFeedist{}

func NewMsgRegisterFeedist(creator string, contract string, shares sdk.Dec) *MsgRegisterFeedist {
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

func validateShares(i interface{}) error {
	v, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("invalid parameter: nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("value cannot be negative: %T", i)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("value cannot be greater than 1: %T", i)
	}

	return nil
}

func (msg *MsgRegisterFeedist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := ethermint.ValidateNonZeroAddress(msg.Contract); err != nil {
		return sdkerrors.Wrapf(err, "invalid contract address %s", msg.Contract)
	}

	if err := validateShares(msg.Shares); err != nil {
		return err
	}

	return nil
}
