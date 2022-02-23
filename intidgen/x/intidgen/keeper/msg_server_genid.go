package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xmonader/intidgen/x/intidgen/types"
)

func (k msgServer) Genid(goCtx context.Context, msg *types.MsgGenid) (*types.MsgGenidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	k.IncTheLastId(ctx)
	return &types.MsgGenidResponse{Id: fmt.Sprint(k.GetTheLastId(ctx))}, nil
}
