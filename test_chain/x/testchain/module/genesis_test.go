package testchain_test

import (
	"testing"

	keepertest "test_chain/testutil/keeper"
	"test_chain/testutil/nullify"
	testchain "test_chain/x/testchain/module"
	"test_chain/x/testchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestchainKeeper(t)
	testchain.InitGenesis(ctx, k, genesisState)
	got := testchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
