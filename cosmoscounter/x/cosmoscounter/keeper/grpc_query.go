package keeper

import (
	"github.com/xmonader/cosmoscounter/x/cosmoscounter/types"
)

var _ types.QueryServer = Keeper{}
