package keeper

import (
	"test_chain/x/testchain/types"
)

var _ types.QueryServer = Keeper{}
