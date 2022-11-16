package keeper

import (
	"context"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) Uploads(goCtx context.Context, req *types.QueryUploadsRequest) (*types.QueryUploadsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of uploads
	var uploads []*types.Upload

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps uploads (using upload key, which is "Upload-value-")
	uploadStore := prefix.NewStore(store, []byte(types.UploadKey))

	// Paginate the posts store based on PageRequest
	pageRes, err := query.Paginate(uploadStore, req.Pagination, func(key []byte, value []byte) error {
		var upload types.Upload
		if err := k.cdc.Unmarshal(value, &upload); err != nil {
			return err
		}

		uploads = append(uploads, &upload)

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of posts and pagination info
	return &types.QueryUploadsResponse{Upload: uploads, Pagination: pageRes}, nil
}

