package feedist

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	feedistsimulation "github.com/evmos/evmos/v9/x/feedist/simulation"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

// avoid unused import issue
var (
	_ = feedistsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterFeedist = "op_weight_msg_register"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterFeedist int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	feedistGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&feedistGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	feedistParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyEnableFeedist), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(feedistParams.EnableFeedist))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterFeedist int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterFeedist, &weightMsgRegisterFeedist, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterFeedist = defaultWeightMsgRegisterFeedist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterFeedist,
		feedistsimulation.SimulateMsgRegisterFeedist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
