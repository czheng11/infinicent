package infinicent_test

import (
	"testing"

	keepertest "github.com/czheng11/infinicent/testutil/keeper"
	"github.com/czheng11/infinicent/testutil/nullify"
	"github.com/czheng11/infinicent/x/infinicent"
	"github.com/czheng11/infinicent/x/infinicent/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InfinicentKeeper(t)
	infinicent.InitGenesis(ctx, *k, genesisState)
	got := infinicent.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
