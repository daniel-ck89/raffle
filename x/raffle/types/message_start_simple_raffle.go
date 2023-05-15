package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartSimpleRaffle = "start_simple_raffle"

var _ sdk.Msg = &MsgStartSimpleRaffle{}

func NewMsgStartSimpleRaffle(creator string, id uint64) *MsgStartSimpleRaffle {
	return &MsgStartSimpleRaffle{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgStartSimpleRaffle) Route() string {
	return RouterKey
}

func (msg *MsgStartSimpleRaffle) Type() string {
	return TypeMsgStartSimpleRaffle
}

func (msg *MsgStartSimpleRaffle) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartSimpleRaffle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartSimpleRaffle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
