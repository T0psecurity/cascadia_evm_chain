package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

// SetFeedist set a specific feedist in the store from its index
func (k Keeper) SetFeedist(ctx sdk.Context, feedist types.Feedist) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedistKeyPrefix))
	b := k.cdc.MustMarshal(&feedist)
	store.Set(types.FeedistKey(
		feedist.Index,
	), b)
}

// GetFeedist returns a feedist from its index
func (k Keeper) GetFeedist(
	ctx sdk.Context,
	index string,

) (val types.Feedist, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedistKeyPrefix))

	b := store.Get(types.FeedistKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeedist removes a feedist from the store
func (k Keeper) RemoveFeedist(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedistKeyPrefix))
	store.Delete(types.FeedistKey(
		index,
	))
}

// GetAllFeedist returns all feedist
func (k Keeper) GetAllFeedist(ctx sdk.Context) (list []types.Feedist) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeedistKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Feedist
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
