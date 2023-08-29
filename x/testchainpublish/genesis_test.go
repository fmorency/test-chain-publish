package testchainpublish_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "test-chain-publish/testutil/keeper"
	"test-chain-publish/testutil/nullify"
	"test-chain-publish/x/testchainpublish"
	"test-chain-publish/x/testchainpublish/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestchainpublishKeeper(t)
	testchainpublish.InitGenesis(ctx, *k, genesisState)
	got := testchainpublish.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
