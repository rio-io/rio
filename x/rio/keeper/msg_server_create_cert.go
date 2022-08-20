package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rio/x/rio/types"
)

func (k msgServer) CreateCert(goCtx context.Context, msg *types.MsgCreateCert) (*types.MsgCreateCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateCertResponse{}, nil
}
