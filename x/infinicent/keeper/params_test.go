package keeper_test

import (
	"testing"

	testkeeper "github.com/czheng11/infinicent/testutil/keeper"
	"github.com/czheng11/infinicent/x/infinicent/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InfinicentKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
