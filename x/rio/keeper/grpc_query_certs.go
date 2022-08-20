package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"rio/x/rio/types"
)

func (k Keeper) Certs(goCtx context.Context, req *types.QueryCertsRequest) (*types.QueryCertsResponse, error) {
	// Throw an error if request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of certs
	var certs []*types.Cert

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps certs (using cert key, which is "Cert-value-")
	certStore := prefix.NewStore(store, []byte(types.CertKey))

	// Paginate the posts store based on PageRequest
	pageRes, err := query.Paginate(certStore, req.Pagination, func(key []byte, value []byte) error {
		var cert types.Cert
		if err := k.cdc.Unmarshal(value, &cert); err != nil {
			return err
		}

		if cert.Owner == req.Address {
			certs = append(certs, &cert)
		}

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of posts and pagination info
	return &types.QueryCertsResponse{Cert: certs, Pagination: pageRes}, nil
}
