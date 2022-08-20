package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rio/x/rio/types"
)

func (k msgServer) SendCert(goCtx context.Context, msg *types.MsgSendCert) (*types.MsgSendCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendCertResponse{}, nil
}
