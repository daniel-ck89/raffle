package keeper

import (
	"context"

	"raffle/x/raffle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Raffles(goCtx context.Context, req *types.QueryRafflesRequest) (*types.QueryRafflesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	queryResult, err := k.GetRaffles(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return queryResult, nil
}
