package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xmonader/cosmos-counter/x/cosmoscounter/types"
)

func (k msgServer) IncrCounter(goCtx context.Context, msg *types.MsgIncrCounter) (*types.MsgIncrCounterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgIncrCounterResponse{}, nil
}
