package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xmonader/intidgen/testutil/keeper"
	"github.com/xmonader/intidgen/x/intidgen/keeper"
	"github.com/xmonader/intidgen/x/intidgen/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IntidgenKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
