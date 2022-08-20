package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rio/x/rio/types"
)

func (k Keeper) Resumes(goCtx context.Context, req *types.QueryResumesRequest) (*types.QueryResumesResponse, error) {
	// Throw an error if request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)

	resume := k.GetResume(ctx, req.Id)

	// Return a struct containing a list of posts and pagination info
	return &types.QueryResumesResponse{Resume: resume}, nil
}
