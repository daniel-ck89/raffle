package keeper

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
	"math/rand"
	"sort"
	"strconv"

	"raffle/x/raffle/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AppendRaffle(ctx sdk.Context, raffle types.Raffle) uint64 {
	count := k.GetRaffleCount(ctx)

	raffle.Id = count
	raffle.Status = 0

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, raffle.Id)

	appendedValue := k.cdc.MustMarshal(&raffle)

	store.Set(byteKey, appendedValue)

	k.SetRaffleCount(ctx, count+1)
	return count
}

func (k Keeper) SimpleRaffle(ctx sdk.Context, msg types.MsgStartSimpleRaffle) error {
	raffle, err := k.GetRaffle(ctx, msg.Id)

	if err != nil {
		return err
	}

	if msg.Creator != raffle.Creator {
		return sdkerrors.Wrapf(types.ErrRaffleFailed,
			"You are not the creator of this raffle.",
		)
	}

	if raffle.Status != 0 {
		return sdkerrors.Wrapf(types.ErrRaffleFailed,
			"This raffle has already exited. Id : %d", msg.Id,
		)
	}

	salt := string(ctx.BlockHeader().AppHash[:]) + strconv.FormatUint(raffle.Id, 10) + raffle.Creator

	loopCnt := uint32(0)
	var winners = make(map[int]string)
	var maxRange = int(raffle.NumberOfParticipants)
	for uint32(len(winners)) < raffle.NumberOfWinners {
		loopCnt++

		if loopCnt > raffle.NumberOfParticipants * 3 {
			return sdkerrors.Wrapf(types.ErrRaffleFailed,
				"Can't handle this rapple. Too many draws. loopCnt : %d", loopCnt,
			)
		}
		h := sha256.New()

		io.WriteString(h, salt + strconv.FormatUint(uint64(loopCnt), 10))
		var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))

		rand.Seed(int64(seed))

		winner := rand.Intn(maxRange)
		rand.Uint64()

		_, exists := winners[winner]
		if !exists {

			winners[winner] = "win"
		}
	}

	sortedWinners, err := getSortedKey(winners)
	if err != nil {
		return err
	}

	raffleResult := types.RaffleResult{
		Id:     msg.Id,
		Result: sortedWinners,
	}

	k.UpdateWinners(ctx, raffleResult)

	k.UpdateRaffleStatus(ctx, raffle)

	return nil
}

func (k Keeper) UpdateWinners(ctx sdk.Context, result types.RaffleResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleResultKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, result.Id)

	raffleResult := k.cdc.MustMarshal(&result)

	store.Set(byteKey, raffleResult)
}

func (k Keeper) UpdateRaffleStatus(ctx sdk.Context, raffle types.Raffle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleKey))

	raffle.Status = 1

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, raffle.Id)

	appendedValue := k.cdc.MustMarshal(&raffle)

	store.Set(byteKey, appendedValue)
}

func (k Keeper) GetRaffle(ctx sdk.Context, id uint64) (types.Raffle, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, id)

	bz := store.Get(byteKey)
	var raffle types.Raffle

	if bz == nil {
		return raffle, sdkerrors.Wrapf(types.ErrRaffleNotFound,
			"Rapple doesn't exist. Id : %d", id,
		)
	}

	err := k.cdc.Unmarshal(bz, &raffle)

	return raffle, err
}

func (k Keeper) GetRaffles(ctx sdk.Context, req *types.QueryRafflesRequest) (*types.QueryRafflesResponse, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleKey))

	var raffles []*types.Raffle

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var raffle types.Raffle
		if err := k.cdc.Unmarshal(value, &raffle); err != nil {
			return err
		}

		raffles = append(raffles, &raffle)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRafflesResponse{Raffles: raffles, Pagination: pageRes}, nil
}

func (k Keeper) GetRaffleResult(ctx sdk.Context, id uint64) (types.RaffleResult, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleResultKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, id)

	bz := store.Get(byteKey)
	var raffleResult types.RaffleResult

	if bz == nil {
		return raffleResult, sdkerrors.Wrapf(types.ErrRaffleNotFound,
			"Rapple result doesn't exist. Id : %d", id,
		)
	}

	err := k.cdc.Unmarshal(bz, &raffleResult)

	return raffleResult, err
}

func (k Keeper) GetRaffleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleCountKey))

	byteKey := []byte(types.RaffleCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetRaffleCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RaffleCountKey))

	byteKey := []byte(types.RaffleCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func getSortedKey(winners map[int]string) (sortedKeys []int32, err error) {

	if len(winners) <= 0 {
		return nil, errors.New("There are no winners.")
	}

	keys := make([]int32, 0, len(winners))
	for k := range winners {
		keys = append(keys, int32(k))
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	return keys, nil

}
