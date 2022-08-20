package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rio/x/rio/types"
)

func (k msgServer) SendCert(goCtx context.Context, msg *types.MsgSendCert) (*types.MsgSendCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Cert
	var cert = types.Cert{
		Creator:     msg.Creator,
		Owner:       msg.To,
		CertType:    msg.CertType,
		Title:       msg.Title,
		Description: msg.Description,
		CreatedAt:   ctx.BlockHeight(),
	}

	// Add a certification to the store and get back the ID
	id := k.AppendCert(ctx, cert)

	// Return the ID of the certrification
	return &types.MsgSendCertResponse{Id: id}, nil
}
