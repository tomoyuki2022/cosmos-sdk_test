package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "test_chain/testutil/keeper"
    "test_chain/x/testchain/types"
    "test_chain/x/testchain/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.TestchainKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}