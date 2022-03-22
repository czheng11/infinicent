package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/czheng11/infinicent/testutil/keeper"
	"github.com/czheng11/infinicent/x/infinicent/keeper"
	"github.com/czheng11/infinicent/x/infinicent/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InfinicentKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
