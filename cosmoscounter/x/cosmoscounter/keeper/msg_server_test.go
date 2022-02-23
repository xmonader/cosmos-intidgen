package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xmonader/cosmoscounter/testutil/keeper"
	"github.com/xmonader/cosmoscounter/x/cosmoscounter/keeper"
	"github.com/xmonader/cosmoscounter/x/cosmoscounter/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmoscounterKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
