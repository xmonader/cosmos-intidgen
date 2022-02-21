package keeper

import (
	"github.com/xmonader/cosmos-counter/x/cosmoscounter/types"
)

var _ types.QueryServer = Keeper{}
