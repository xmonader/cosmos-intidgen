package cosmoscounter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/xmonader/cosmoscounter/testutil/keeper"
	"github.com/xmonader/cosmoscounter/testutil/nullify"
	"github.com/xmonader/cosmoscounter/x/cosmoscounter"
	"github.com/xmonader/cosmoscounter/x/cosmoscounter/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoscounterKeeper(t)
	cosmoscounter.InitGenesis(ctx, *k, genesisState)
	got := cosmoscounter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
