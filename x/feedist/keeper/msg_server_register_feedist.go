package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

func (k msgServer) RegisterFeedist(goCtx context.Context, msg *types.MsgRegisterFeedist) (*types.MsgRegisterFeedistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterFeedistResponse{}, nil
}
