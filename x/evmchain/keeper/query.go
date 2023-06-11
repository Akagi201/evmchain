package keeper

import (
	"github.com/Akagi201/evmchain/x/evmchain/types"
)

var _ types.QueryServer = Keeper{}
