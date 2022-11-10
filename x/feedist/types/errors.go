package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/feedist module sentinel errors
var (
	ErrSample                    = sdkerrors.Register(ModuleName, 0, "sample error")
	ErrFeedistEnable             = sdkerrors.Register(ModuleName, 1, "Feedist module is disabled")
	ErrRevenueNoContractDeployed = sdkerrors.Register(ModuleName, 2, "This contract is not deployed")
	ErrUnauthorized              = sdkerrors.Register(ModuleName, 3, "Unauthorized")
)
