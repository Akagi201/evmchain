package keeper_test

import (
	"testing"

	testkeeper "github.com/Akagi201/evmchain/testutil/keeper"
	"github.com/Akagi201/evmchain/x/evmchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EvmchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
