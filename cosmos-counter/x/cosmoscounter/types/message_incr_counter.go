package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIncrCounter = "incr_counter"

var _ sdk.Msg = &MsgIncrCounter{}

func NewMsgIncrCounter(creator string) *MsgIncrCounter {
	return &MsgIncrCounter{
		Creator: creator,
	}
}

func (msg *MsgIncrCounter) Route() string {
	return RouterKey
}

func (msg *MsgIncrCounter) Type() string {
	return TypeMsgIncrCounter
}

func (msg *MsgIncrCounter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIncrCounter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIncrCounter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
