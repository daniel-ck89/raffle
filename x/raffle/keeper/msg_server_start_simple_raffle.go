package keeper

import (
	"context"

	"raffle/x/raffle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StartSimpleRaffle(goCtx context.Context, msg *types.MsgStartSimpleRaffle) (*types.MsgStartSimpleRaffleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.SimpleRaffle(ctx, *msg)

	if err != nil {
		return nil, err
	}

	return &types.MsgStartSimpleRaffleResponse{}, nil
}
