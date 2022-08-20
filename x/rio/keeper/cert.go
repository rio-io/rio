package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"rio/x/rio/types"
)

// Get the correct number of certs in the store
func (k Keeper) GetCertCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "rio") and CertCountKey (which is "Cert-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CertCountKey))

	// Convert the CertCountKey to bytes
	byteKey := []byte(types.CertCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first cert)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// Set the certification count
func (k Keeper) SetCertCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "rio") and CertCountKey (which is "Cert-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CertCountKey))

	// Convert the CertCountKey to bytes
	byteKey := []byte(types.CertCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Cert-count- to count
	store.Set(byteKey, bz)
}

// Add a certication to the store and get back the ID
func (k Keeper) AppendCert(ctx sdk.Context, cert types.Cert) uint64 {
	// Get the current number of posts in the store (count)
	count := k.GetCertCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	cert.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CertKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, cert.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&cert)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetCertCount(ctx, count+1)
	return count
}
