package intidgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/xmonader/intidgen/testutil/keeper"
	"github.com/xmonader/intidgen/testutil/nullify"
	"github.com/xmonader/intidgen/x/intidgen"
	"github.com/xmonader/intidgen/x/intidgen/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IntidgenKeeper(t)
	intidgen.InitGenesis(ctx, *k, genesisState)
	got := intidgen.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
