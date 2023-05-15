package keeper

import (
	"context"

	"raffle/x/raffle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSimpleRaffle(goCtx context.Context, msg *types.MsgCreateSimpleRaffle) (*types.MsgCreateSimpleRaffleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	raffle := types.Raffle{
		Creator:              msg.Creator,
		Title:                msg.Title,
		Description:          msg.Description,
		ParticipantListUrl:   msg.ParticipantListUrl,
		NumberOfWinners:      msg.NumberOfWinners,
		NumberOfParticipants: msg.NumberOfParticipants,
	}

	raffleId := k.AppendRaffle(ctx, raffle)

	return &types.MsgCreateSimpleRaffleResponse{Id: raffleId}, nil
}
