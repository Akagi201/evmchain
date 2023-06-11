package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Akagi201/evmchain/testutil/keeper"
	"github.com/Akagi201/evmchain/x/evmchain/keeper"
	"github.com/Akagi201/evmchain/x/evmchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EvmchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
