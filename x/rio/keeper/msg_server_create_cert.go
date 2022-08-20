package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rio/x/rio/types"
)

func (k msgServer) CreateCert(goCtx context.Context, msg *types.MsgCreateCert) (*types.MsgCreateCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Cert
	var cert = types.Cert {
		Creator: msg.Creator,
		Title: msg.Title,
	}

	// Add a certification to the store and get back the ID
	id := k.AppendCert(ctx, cert)

	// Return the ID of the certrification
	return &types.MsgCreateCertResponse{Id: id}, nil
}
