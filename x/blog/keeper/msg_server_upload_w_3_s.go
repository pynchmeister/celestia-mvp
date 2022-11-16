package keeper

import (
	"context"
	"fmt"
	"os"

	"github.com/web3-storage/go-w3s-client"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UploadW3S(goCtx context.Context, msg *types.MsgUploadW3S) (*types.MsgUploadW3SResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	c, err := w3s.NewClient(w3s.WithToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDE3ZkZFQ0Q2OGU3MDM5NTNmRURlNDZjMTE2ZUFiN0NEQTVERjZDNmIiLCJpc3MiOiJ3ZWIzLXN0b3JhZ2UiLCJpYXQiOjE2Njg2MDIzMTMyMTIsIm5hbWUiOiJ0ZXN0In0.NmkFzUGexxOAgKVxOKdZhfA8ZfAW-0ftHJ5p_wQaUtY"))
	if err != nil {
		panic(err)
	}

	// random file access and upload successful
	f, err := os.Open("exampleDir/IMG_4624.jpg")

	if err != nil {
		panic(err)
	}

	cid, err := c.Put(ctx, f)
	if err != nil {
		panic(err)
	}

	fmt.Println(cid)

	// Create variable of type Post
	var upload = types.Upload{
		Creator: msg.Creator,
		Title:   msg.Title,
		Content: msg.Content,
	}

	// Add a post to the store and get back the ID (CID?)
	id := k.AppendUpload(ctx, upload)

	// TODO: Handling the message
	// _ = ctx

	return &types.MsgUploadW3SResponse{Id: id}, nil
}