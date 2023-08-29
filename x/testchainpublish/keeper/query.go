package keeper

import (
	"test-chain-publish/x/testchainpublish/types"
)

var _ types.QueryServer = Keeper{}
