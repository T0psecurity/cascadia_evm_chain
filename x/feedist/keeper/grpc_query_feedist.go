package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FeedistAll(c context.Context, req *types.QueryAllFeedistRequest) (*types.QueryAllFeedistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var feedists []types.Feedist
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	feedistStore := prefix.NewStore(store, types.KeyPrefix(types.FeedistKeyPrefix))

	pageRes, err := query.Paginate(feedistStore, req.Pagination, func(key []byte, value []byte) error {
		var feedist types.Feedist
		if err := k.cdc.Unmarshal(value, &feedist); err != nil {
			return err
		}

		feedists = append(feedists, feedist)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeedistResponse{Feedist: feedists, Pagination: pageRes}, nil
}

func (k Keeper) Feedist(c context.Context, req *types.QueryGetFeedistRequest) (*types.QueryGetFeedistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFeedist(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFeedistResponse{Feedist: val}, nil
}
