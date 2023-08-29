package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "test-chain-publish/testutil/keeper"
	"test-chain-publish/x/testchainpublish/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestchainpublishKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
