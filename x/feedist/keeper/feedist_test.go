package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/evmos/evmos/v9/testutil/keeper"
	"github.com/evmos/evmos/v9/testutil/nullify"
	"github.com/evmos/evmos/v9/x/feedist/keeper"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFeedist(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Feedist {
	items := make([]types.Feedist, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFeedist(ctx, items[i])
	}
	return items
}

func TestFeedistGet(t *testing.T) {
	keeper, ctx := keepertest.FeedistKeeper(t)
	items := createNFeedist(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFeedist(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFeedistRemove(t *testing.T) {
	keeper, ctx := keepertest.FeedistKeeper(t)
	items := createNFeedist(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeedist(ctx,
			item.Index,
		)
		_, found := keeper.GetFeedist(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFeedistGetAll(t *testing.T) {
	keeper, ctx := keepertest.FeedistKeeper(t)
	items := createNFeedist(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeedist(ctx)),
	)
}
