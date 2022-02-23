package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/xmonader/intidgen/testutil/keeper"
	"github.com/xmonader/intidgen/x/intidgen/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IntidgenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
