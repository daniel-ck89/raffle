package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/raffle module sentinel errors
var (
	ErrRaffleTitleTooLarge                                = sdkerrors.Register(ModuleName, 1100, "title too large")
	ErrRaffleDescriptionTooLarge                          = sdkerrors.Register(ModuleName, 1101, "description too large")
	ErrRaffleParticipantListUrlTooLarge                   = sdkerrors.Register(ModuleName, 1102, "participantListUrl too large")
	ErrRaffleNumberOfWinnersOutOfRange                    = sdkerrors.Register(ModuleName, 1103, "NumberOfWinners out of range")
	ErrRaffleNumberOfParticipantOutOfRange                = sdkerrors.Register(ModuleName, 1104, "NumberOfParticipants out of range")
	ErrRaffleNumberOfWinnerGreaterThanNumberOfParticipant = sdkerrors.Register(ModuleName, 1105, "NumberOfWinner is greater than NumberOfParticipant")
	ErrRaffleNotFound                                     = sdkerrors.Register(ModuleName, 1106, "raffle not found")
	ErrRaffleFailed                                       = sdkerrors.Register(ModuleName, 1107, "raffle failed")
	ErrRaffleResultNotFound                               = sdkerrors.Register(ModuleName, 1108, "raffle result not found")
)
