package rio_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "rio/testutil/keeper"
	"rio/testutil/nullify"
	"rio/x/rio"
	"rio/x/rio/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RioKeeper(t)
	rio.InitGenesis(ctx, *k, genesisState)
	got := rio.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
