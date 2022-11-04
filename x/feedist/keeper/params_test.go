package keeper_test

import (
	"testing"

	testkeeper "github.com/evmos/evmos/v9/testutil/keeper"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FeedistKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.EnableFeedist, k.EnableFeedist(ctx))
}
