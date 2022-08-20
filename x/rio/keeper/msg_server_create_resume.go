package keeper

import (
	"context"

	"rio/x/rio/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateResume(goCtx context.Context, msg *types.MsgCreateResume) (*types.MsgCreateResumeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	certs, _ := k.GetCerts(ctx, msg.Certs)

	// Create variable of type Resume
	var resume = types.Resume{
		Creator:     msg.Creator,
		AvatarUrl:   msg.AvatarUrl,
		Name:        msg.Name,
		Description: msg.Description,
		Certs:       certs,
		CreatedAt:   ctx.BlockHeight(),
	}

	id := k.AppendResume(ctx, resume)

	return &types.MsgCreateResumeResponse{Id: id}, nil
}
