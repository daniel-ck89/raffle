package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSimpleRaffle  = "create_simple_raffle"
	MAX_NUNBER_OF_WINNERS      = 10000
	MAX_NUNBER_OF_PARTICIPANTS = 30000
)

var _ sdk.Msg = &MsgCreateSimpleRaffle{}

func NewMsgCreateSimpleRaffle(creator string, title string, description string, participantListUrl string, numberOfWinners uint32, numberOfParticipants uint32) *MsgCreateSimpleRaffle {
	return &MsgCreateSimpleRaffle{
		Creator:              creator,
		Title:                title,
		Description:          description,
		ParticipantListUrl:   participantListUrl,
		NumberOfWinners:      numberOfWinners,
		NumberOfParticipants: numberOfParticipants,
	}
}

func (msg *MsgCreateSimpleRaffle) Route() string {
	return RouterKey
}

func (msg *MsgCreateSimpleRaffle) Type() string {
	return TypeMsgCreateSimpleRaffle
}

func (msg *MsgCreateSimpleRaffle) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSimpleRaffle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSimpleRaffle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	titleLength := len(msg.Title)
	if uint64(titleLength) > 256 {
		return sdkerrors.Wrapf(ErrRaffleTitleTooLarge,
			"maximum number of characters is %d but received %d characters",
			256, titleLength,
		)
	}

	descriptionLength := len(msg.Description)
	if uint64(descriptionLength) > 1024 {
		return sdkerrors.Wrapf(ErrRaffleDescriptionTooLarge,
			"maximum number of characters is %d but received %d characters",
			1024, descriptionLength,
		)
	}

	participantListUrlLength := len(msg.ParticipantListUrl)
	if uint64(participantListUrlLength) > 2083 {
		return sdkerrors.Wrapf(ErrRaffleParticipantListUrlTooLarge,
			"maximum number of characters is %d but received %d characters",
			2083, participantListUrlLength,
		)
	}

	if msg.NumberOfWinners == 0 || msg.NumberOfWinners > MAX_NUNBER_OF_WINNERS {
		return sdkerrors.Wrapf(ErrRaffleNumberOfWinnersOutOfRange,
			"The acceptable range for NumberOfWinners is 1 to %d. but received %d", MAX_NUNBER_OF_WINNERS, msg.NumberOfWinners,
		)
	}

	if msg.NumberOfParticipants == 0 || msg.NumberOfParticipants > MAX_NUNBER_OF_PARTICIPANTS {
		return sdkerrors.Wrapf(ErrRaffleNumberOfParticipantOutOfRange,
			"The acceptable range for NumberOfParticipants is 1 to %d. but received %d", MAX_NUNBER_OF_PARTICIPANTS, msg.NumberOfParticipants,
		)
	}

	if msg.NumberOfWinners > msg.NumberOfParticipants {
		return sdkerrors.Wrapf(ErrRaffleNumberOfWinnerGreaterThanNumberOfParticipant,
			"maximum NumberOfWinners is %d but received %d",
			msg.NumberOfParticipants, msg.NumberOfWinners,
		)
	}

	return nil
}
