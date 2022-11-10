package feedist_test

import (
	"testing"

	keepertest "github.com/evmos/evmos/v9/testutil/keeper"
	"github.com/evmos/evmos/v9/x/feedist"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FeedistList: []types.Feedist{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FeedistKeeper(t)
	feedist.InitGenesis(ctx, *k, genesisState)
	got := feedist.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.FeedistList, got.FeedistList)
	// this line is used by starport scaffolding # genesis/test/assert
}
