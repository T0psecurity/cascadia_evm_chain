package keeper

import (
	"github.com/evmos/evmos/v9/x/feedist/types"
)

var _ types.QueryServer = Keeper{}
