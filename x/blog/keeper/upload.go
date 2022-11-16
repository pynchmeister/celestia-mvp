package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k Keeper) AppendUpload(ctx sdk.Context, upload types.Upload) uint64 {
	// Get the current number of uploads in the store
	count := k.GetUploadCount(ctx)

	// Assign an ID to the upload based on the number of uploads in the store
	upload.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UploadKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, upload.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&upload)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetUploadCount(ctx, count+1)
	return count
}

func (k Keeper) GetUploadCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and UploadCountKey (which is "Upload/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UploadCountKey))

	// Convert the UploadCountKey to bytes
	byteKey := []byte(types.UploadCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetUploadCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and UploadCountKey (which is "Upload/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UploadCountKey))

	// Convert the UploadCountKey to bytes
	byteKey := []byte(types.UploadCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Upload/count/ to count
	store.Set(byteKey, bz)
}

// func (k Keeper) AppendPost() uint64 {
//   count := k.GetPostCount()
//   store.Set()
//   k.SetPostCount()
//   return count
// }
