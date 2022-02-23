package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/xmonader/cosmoscounter/testutil/keeper"
	"github.com/xmonader/cosmoscounter/x/cosmoscounter/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmoscounterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
