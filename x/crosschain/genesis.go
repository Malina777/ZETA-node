package crosschain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/crosschain/keeper"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the outTxTracker
	for _, elem := range genState.OutTxTrackerList {
		k.SetOutTxTracker(ctx, elem)
	}
	// Set all the inTxHashToCctx
	for _, elem := range genState.InTxHashToCctxList {
		k.SetInTxHashToCctx(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	// Set if defined
	if genState.Keygen != nil {
		k.SetKeygen(ctx, *genState.Keygen)
	}

	// Set all the tSSVoter
	for _, elem := range genState.TSSVoterList {
		k.SetTSSVoter(ctx, *elem)
	}

	// Set all the tSS
	for _, elem := range genState.TSSList {
		k.SetTSS(ctx, *elem)
	}

	// Set all the gasBalance
	for _, elem := range genState.GasBalanceList {
		k.SetGasBalance(ctx, *elem)
	}

	// Set all the gasPrice
	for _, elem := range genState.GasPriceList {
		k.SetGasPrice(ctx, *elem)
	}

	// Set all the chainNonces
	for _, elem := range genState.ChainNoncesList {
		k.SetChainNonces(ctx, *elem)
	}

	// Set all the lastBlockHeight
	for _, elem := range genState.LastBlockHeightList {
		k.SetLastBlockHeight(ctx, *elem)
	}

	// Set all the send
	for _, elem := range genState.CrossChainTxs {
		k.SetCrossChainTx(ctx, *elem)
	}

	// Set all the nodeAccount
	for _, elem := range genState.NodeAccountList {
		k.SetNodeAccount(ctx, *elem)
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.OutTxTrackerList = k.GetAllOutTxTracker(ctx)
	genesis.InTxHashToCctxList = k.GetAllInTxHashToCctx(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	// Get all keygen
	keygen, found := k.GetKeygen(ctx)
	if found {
		genesis.Keygen = &keygen
	}

	// Get all tSSVoter
	tSSVoterList := k.GetAllTSSVoter(ctx)
	for _, elem := range tSSVoterList {
		elem := elem
		genesis.TSSVoterList = append(genesis.TSSVoterList, &elem)
	}

	// Get all tSS
	tSSList := k.GetAllTSS(ctx)
	for _, elem := range tSSList {
		elem := elem
		genesis.TSSList = append(genesis.TSSList, &elem)
	}

	// Get all gasBalance
	gasBalanceList := k.GetAllGasBalance(ctx)
	for _, elem := range gasBalanceList {
		elem := elem
		genesis.GasBalanceList = append(genesis.GasBalanceList, &elem)
	}

	// Get all gasPrice
	gasPriceList := k.GetAllGasPrice(ctx)
	for _, elem := range gasPriceList {
		elem := elem
		genesis.GasPriceList = append(genesis.GasPriceList, &elem)
	}

	// Get all chainNonces
	chainNoncesList := k.GetAllChainNonces(ctx)
	for _, elem := range chainNoncesList {
		elem := elem
		genesis.ChainNoncesList = append(genesis.ChainNoncesList, &elem)
	}

	// Get all lastBlockHeight
	lastBlockHeightList := k.GetAllLastBlockHeight(ctx)
	for _, elem := range lastBlockHeightList {
		elem := elem
		genesis.LastBlockHeightList = append(genesis.LastBlockHeightList, &elem)
	}

	// Get all send
	sendList := k.GetAllCctxByStatuses(ctx, types.AllStatus())
	for _, elem := range sendList {
		elem := elem
		genesis.CrossChainTxs = append(genesis.CrossChainTxs, elem)
	}

	// Get all nodeAccount
	nodeAccountList := k.GetAllNodeAccount(ctx)
	for _, elem := range nodeAccountList {
		elem := elem
		genesis.NodeAccountList = append(genesis.NodeAccountList, &elem)
	}
	return genesis
}

// TODO : Verify genesis import and export
