package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	EVMConfig(ctx sdk.Context) (*evmtypes.EVMConfig, error)
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
}
