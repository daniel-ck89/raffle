package keeper

import (
	"context"
	"encoding/json"

	"raffle/x/raffle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RaffleResult(goCtx context.Context, req *types.QueryRaffleResultRequest) (*types.QueryRaffleResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	raffleResult, err := k.GetRaffleResult(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(raffleResult)

	if err != nil {
		return nil, err
	}

	return &types.QueryRaffleResultResponse{RaffleResult: string(result)}, nil
}
