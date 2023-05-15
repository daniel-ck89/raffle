package keeper

import (
	"raffle/x/raffle/types"
)

var _ types.QueryServer = Keeper{}
