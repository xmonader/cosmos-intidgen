package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/xmonader/intidgen/x/intidgen/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		lastid     uint64
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetThelastid set thelastid in the store
func (k Keeper) IncTheLastId(ctx sdk.Context) {
	newVal := k.GetTheLastId(ctx) + 1
	k.Logger(ctx).Info(">>>>>>> IncTheLastId", "newVal", newVal)
	k.SetTheLastId(ctx, newVal)
}

// SetThelastid set thelastid in the store
func (k Keeper) SetTheLastId(ctx sdk.Context, to uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix("lastid-value-"))
	k.Logger(ctx).Info(">>>>> SetTheLastId", "to", to)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, to)
	store.Set([]byte{0}, bz)
}

// GetThelastid returns thelastid
func (k Keeper) GetTheLastId(ctx sdk.Context) (lastId uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix("lastid-value-"))

	b := store.Get([]byte{0})

	if b == nil {
		k.Logger(ctx).Info(">>>>> GetTheLastId", "lastId b doesn't exist", b)
		k.SetTheLastId(ctx, 1)
		return uint64(1)
	}
	k.Logger(ctx).Info(">>>>> GetTheLastId", "lastId b exists", b)
	lastId = binary.BigEndian.Uint64(b)
	return lastId

}
