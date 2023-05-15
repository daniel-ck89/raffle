package keeper

import (
	"context"

	"raffle/x/raffle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Raffle(goCtx context.Context, req *types.QueryRaffleRequest) (*types.QueryRaffleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	raffle, err := k.GetRaffle(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &types.QueryRaffleResponse{Raffle: &raffle}, nil
}
