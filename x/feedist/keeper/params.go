package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.EnableFeedist(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// EnableFeedist returns the EnableFeedist param
func (k Keeper) EnableFeedist(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyEnableFeedist, &res)
	return
}
