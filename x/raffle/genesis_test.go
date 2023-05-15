package raffle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "raffle/testutil/keeper"
	"raffle/testutil/nullify"
	"raffle/x/raffle"
	"raffle/x/raffle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RaffleKeeper(t)
	raffle.InitGenesis(ctx, *k, genesisState)
	got := raffle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
