package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"rio/x/rio/types"
)

// Get the correct number of resume in the store
func (k Keeper) GetResumeCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "rio") and ResumeCountKey (which is "Resume-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ResumeCountKey))

	// Convert the ResumeCountKey to bytes
	byteKey := []byte(types.ResumeCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first resume)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// Set the resume count
func (k Keeper) SetResumeCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "rio") and ResumeCountKey (which is "Resume-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ResumeCountKey))

	// Convert the ResumeCountKey to bytes
	byteKey := []byte(types.ResumeCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Resume-count- to count
	store.Set(byteKey, bz)
}

// Add a resum to the store and get back the ID
func (k Keeper) AppendResume(ctx sdk.Context, resume types.Resume) uint64 {
	// Get the current number of posts in the store (count)
	count := k.GetResumeCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	resume.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ResumeKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, resume.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&resume)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetResumeCount(ctx, count+1)
	return count
}