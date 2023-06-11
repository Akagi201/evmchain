package evmchain_test

import (
	"testing"

	keepertest "github.com/Akagi201/evmchain/testutil/keeper"
	"github.com/Akagi201/evmchain/testutil/nullify"
	"github.com/Akagi201/evmchain/x/evmchain"
	"github.com/Akagi201/evmchain/x/evmchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EvmchainKeeper(t)
	evmchain.InitGenesis(ctx, *k, genesisState)
	got := evmchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
