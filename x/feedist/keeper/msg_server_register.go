package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

func (k msgServer) RegisterFeedist(goCtx context.Context, msg *types.MsgRegisterFeedist) (*types.MsgRegisterFeedistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	params := k.GetParams(ctx)
	if !params.EnableFeedist {
		return nil, types.ErrFeedistEnable
	}

	if msg.Creator != "cascadia1fv3uch395zu60kmme0yg0tsrgrsgnjvfn7fwn4" {
		return nil, types.ErrUnauthorized
	}

	contract := common.HexToAddress(msg.Contract)

	// contract must already be deployed, to avoid spam registrations
	contractAccount := k.evmKeeper.GetAccountWithoutBalance(ctx, contract)

	if contractAccount == nil || !contractAccount.IsContract() {
		return nil, sdkerrors.Wrapf(
			types.ErrRevenueNoContractDeployed,
			"no contract code found at address %s", msg.Contract,
		)
	}

	k.SetFeedist(ctx, types.Feedist{
		Index:    "feedist",
		Contract: msg.Contract,
		Shares:   msg.Shares,
	})

	return &types.MsgRegisterFeedistResponse{}, nil
}
