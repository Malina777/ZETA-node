package main

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/observer/types"
	"strconv"
	"strings"
)

func AddObserverAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-observer-list [list of observers]",
		Short: "Add a list of observers to the observer mapper",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			observerList, err := types.ParsefileToObserverMapper(args[0])
			if err != nil {
				return err
			}
			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			zetaObserverGenState := types.GetGenesisStateFromAppState(cdc, appState)
			zetaObserverGenState.Observers = observerList

			zetaObserverStateBz, err := json.Marshal(zetaObserverGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal Observer List into Genesis File: %w", err)
			}
			appState[types.ModuleName] = zetaObserverStateBz
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}
	return cmd
}

func AddObserverAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-observer [chainName] [chainID][comma separate list of address] ",
		Short: "Add a list of observers to the observer mapper",
		Long: `
           Chain Types :
					"Empty"     
					"Eth"       
					"ZetaChain" 
					"Btc"       
					"Polygon"   
					"BscMainnet"
					"Goerli"    
					"Mumbai"    
					"Ropsten"   
					"Ganache"   
					"Baobap"    
					"BscTestnet"
					"BTCTestnet"
			`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			chainName := common.ParseChainName(args[0])
			chainID, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			observer := &types.ObserverMapper{
				Index: "",
				ObserverChain: &common.Chain{
					ChainName: chainName,
					ChainId:   int64(chainID),
				},
				ObserverList: strings.Split(args[3], ","),
			}
			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			zetaObserverGenState := types.GetGenesisStateFromAppState(cdc, appState)
			zetaObserverGenState.Observers = append(zetaObserverGenState.Observers, observer)

			zetaObserverStateBz, err := json.Marshal(zetaObserverGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal Observer List into Genesis File: %w", err)
			}
			appState[types.ModuleName] = zetaObserverStateBz
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}
	return cmd
}
