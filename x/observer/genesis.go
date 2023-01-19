package observer

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/observer/keeper"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	fmt.Println("Genesis Account : ")
	genesisObservers := genState.Observers
	fmt.Println("genesisObservers : ", genesisObservers)
	types.VerifyObserverMapper(genesisObservers)
	for _, mapper := range genesisObservers {
		fmt.Println("SetObserverMapper : ", mapper)
		k.SetObserverMapper(ctx, mapper)
	}
	k.SetParams(ctx, types.DefaultParams())
	k.SetSupportedChain(ctx, types.SupportedChains{ChainList: common.DefaultChainsList()})
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	return &types.GenesisState{
		Ballots:   k.GetAllBallots(ctx),
		Observers: k.GetAllObserverMappers(ctx),
		Params:    &params,
	}
}
