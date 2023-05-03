package main

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/zetacore/app"
	"github.com/zeta-chain/zetacore/cmd/zetacored/config"
	"github.com/zeta-chain/zetacore/common"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	"github.com/zeta-chain/zetacore/x/observer/types"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ObserverBalance = "100000000000000000000000"
	HotkeyBalance   = "100000000000000000000"
)

func AddObserverAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-observer-list [observer-list.json] ",
		Short: "Add a list of observers to the observer mapper ,default path is ~/.zetacored/os_info/observer_info.json",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec
			serverCtx := server.GetServerContextFromCmd(cmd)
			serverConfig := serverCtx.Config

			defaultHome := app.DefaultNodeHome
			defaultFile := filepath.Join(defaultHome, "os_info", "observer_info.json")
			if len(args) == 0 {
				args = append(args, defaultFile)
			}
			file := args[0]
			observerInfo, err := ParsefileToObserverDetails(file)
			if err != nil {
				return err
			}
			var observerMapper []*types.ObserverMapper
			var grantAuthorizations []authz.GrantAuthorization
			var nodeAccounts []*crosschaintypes.NodeAccount
			observersForChain := map[int64][]string{}
			// DefaultChainsList is based on Build Flags
			supportedChains := common.DefaultChainsList()
			var balances []banktypes.Balance
			commonCoins, ok := sdk.NewIntFromString(ObserverBalance)
			if !ok {
				panic("Failed to parse string to int for observer")
			}
			commonHotkeyCoins, ok := sdk.NewIntFromString(HotkeyBalance)
			if !ok {
				panic("Failed to parse string to int for hotkey")
			}
			commonObserverBalance := sdk.NewCoins(sdk.NewCoin(config.BaseDenom, commonCoins))
			commonHotkeyBalance := sdk.NewCoins(sdk.NewCoin(config.BaseDenom, commonHotkeyCoins))
			// Generate the grant authorizations and created observer list for chain
			for _, info := range observerInfo {

				balances = append(balances, banktypes.Balance{
					Address: info.ObserverAddress,
					Coins:   commonObserverBalance,
				})
				if isValidatorOnly(info.IsObserver) {
					continue
				}

				if info.ZetaClientGranteeAddress == "" || info.ObserverAddress == "" {
					panic("ZetaClientGranteeAddress or ObserverAddress is empty")
				}
				grantAuthorizations = append(grantAuthorizations, generateGrants(info)...)
				for _, chain := range supportedChains {
					observersForChain[chain.ChainId] = append(observersForChain[chain.ChainId], info.ObserverAddress)
				}
				if info.ZetaClientGranteePubKey != "" {
					pubkey, err := common.NewPubKey(info.ZetaClientGranteePubKey)
					if err != nil {
						panic(err)
					}
					pubkeySet := common.PubKeySet{
						Secp256k1: pubkey,
						Ed25519:   "",
					}
					na := crosschaintypes.NodeAccount{
						Creator:          info.ObserverAddress,
						TssSignerAddress: info.ZetaClientGranteeAddress,
						PubkeySet:        &pubkeySet,
						NodeStatus:       crosschaintypes.NodeStatus_Active,
					}
					nodeAccounts = append(nodeAccounts, &na)
				}

				balances = append(balances, banktypes.Balance{
					Address: info.ZetaClientGranteeAddress,
					Coins:   commonHotkeyBalance,
				})
			}

			// Generate observer mappers for each chain
			for chainID, observers := range observersForChain {
				observers = removeDuplicate(observers)
				chain := common.GetChainFromChainID(chainID)
				mapper := types.ObserverMapper{
					ObserverChain: chain,
					ObserverList:  observers,
				}
				observerMapper = append(observerMapper, &mapper)
			}

			genFile := serverConfig.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			// Add node accounts to cross chain genesis state
			zetaCrossChainGenState := crosschaintypes.GetGenesisStateFromAppState(cdc, appState)
			zetaCrossChainGenState.NodeAccountList = nodeAccounts

			// Add observers to observer genesis state
			zetaObserverGenState := types.GetGenesisStateFromAppState(cdc, appState)
			zetaObserverGenState.Observers = observerMapper

			// Add grant authorizations to authz genesis state
			var authzGenState authz.GenesisState
			if appState[authz.ModuleName] != nil {
				err := cdc.UnmarshalJSON(appState[authz.ModuleName], &authzGenState)
				if err != nil {
					panic(fmt.Sprintf("Failed to get genesis state from app state: %s", err.Error()))
				}
			}

			authzGenState.Authorization = grantAuthorizations

			// Marshal modified states into genesis file
			zetaCrossChainStateBz, err := json.Marshal(zetaCrossChainGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal Observer List into Genesis File: %w", err)
			}
			zetaObserverStateBz, err := json.Marshal(zetaObserverGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal Observer List into Genesis File: %w", err)
			}
			err = codectypes.UnpackInterfaces(authzGenState, cdc)
			if err != nil {
				return fmt.Errorf("failed to authz grants into upackeder: %w", err)
			}
			authZStateBz, err := cdc.MarshalJSON(&authzGenState)
			if err != nil {
				return fmt.Errorf("failed to authz grants into Genesis File: %w", err)
			}
			appState[types.ModuleName] = zetaObserverStateBz
			appState[authz.ModuleName] = authZStateBz
			appState[crosschaintypes.ModuleName] = zetaCrossChainStateBz
			modifiedAppState, err := AddGenesisAccount(clientCtx, balances, appState)
			if err != nil {
				panic(err)
			}
			// Create new genesis file
			appStateJSON, err := json.Marshal(modifiedAppState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}
	return cmd
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func generateGrants(info ObserverInfoReader) []authz.GrantAuthorization {
	sdk.MustAccAddressFromBech32(info.ObserverAddress)
	var grants []authz.GrantAuthorization
	if info.ZetaClientGranteeAddress != "" {
		sdk.MustAccAddressFromBech32(info.ZetaClientGranteeAddress)
		grants = append(grants, addZetaClientGrants(grants, info)...)
	}
	if info.SpendGranteeAddress != "" {
		sdk.MustAccAddressFromBech32(info.SpendGranteeAddress)
		grants = append(grants, addSpendingGrants(grants, info)...)
	}
	if info.StakingGranteeAddress != "" {
		sdk.MustAccAddressFromBech32(info.StakingGranteeAddress)
		grants = append(grants, addStakingGrants(grants, info)...)
	}

	if info.GovGranteeAddress != "" {
		sdk.MustAccAddressFromBech32(info.GovGranteeAddress)
		grants = append(grants, addGovGrants(grants, info)...)
	}

	return grants
}

func addZetaClientGrants(grants []authz.GrantAuthorization, info ObserverInfoReader) []authz.GrantAuthorization {
	txTypes := crosschaintypes.GetAllAuthzZetaclientTxTypes()
	for _, txType := range txTypes {
		auth, err := codectypes.NewAnyWithValue(authz.NewGenericAuthorization(txType))
		if err != nil {
			panic(err)
		}
		grants = append(grants, authz.GrantAuthorization{
			Granter:       info.ObserverAddress,
			Grantee:       info.ZetaClientGranteeAddress,
			Authorization: auth,
			Expiration:    nil,
		})
	}
	return grants
}

func addGovGrants(grants []authz.GrantAuthorization, info ObserverInfoReader) []authz.GrantAuthorization {

	txTypes := []string{sdk.MsgTypeURL(&v1beta1.MsgVote{}),
		sdk.MsgTypeURL(&v1beta1.MsgSubmitProposal{}),
		sdk.MsgTypeURL(&v1beta1.MsgDeposit{}),
		sdk.MsgTypeURL(&v1beta1.MsgVoteWeighted{}),
		sdk.MsgTypeURL(&v1.MsgVote{}),
		sdk.MsgTypeURL(&v1.MsgSubmitProposal{}),
		sdk.MsgTypeURL(&v1.MsgDeposit{}),
		sdk.MsgTypeURL(&v1.MsgVoteWeighted{}),
	}
	for _, txType := range txTypes {
		auth, err := codectypes.NewAnyWithValue(authz.NewGenericAuthorization(txType))
		if err != nil {
			panic(err)
		}
		grants = append(grants, authz.GrantAuthorization{
			Granter:       info.ObserverAddress,
			Grantee:       info.GovGranteeAddress,
			Authorization: auth,
			Expiration:    nil,
		})
	}
	return grants
}

func addSpendingGrants(grants []authz.GrantAuthorization, info ObserverInfoReader) []authz.GrantAuthorization {
	spendMaxTokens, ok := sdk.NewIntFromString(info.SpendMaxTokens)
	if !ok {
		panic("Failed to parse spend max tokens")
	}
	spendAuth, err := codectypes.NewAnyWithValue(&banktypes.SendAuthorization{
		SpendLimit: sdk.NewCoins(sdk.NewCoin(config.BaseDenom, spendMaxTokens)),
	})
	if err != nil {
		panic(err)
	}
	grants = append(grants, authz.GrantAuthorization{
		Granter:       info.ObserverAddress,
		Grantee:       info.SpendGranteeAddress,
		Authorization: spendAuth,
		Expiration:    nil,
	})
	return grants
}

func addStakingGrants(grants []authz.GrantAuthorization, info ObserverInfoReader) []authz.GrantAuthorization {
	stakingMaxTokens, ok := sdk.NewIntFromString(info.StakingMaxTokens)
	if !ok {
		panic("Failed to parse staking max tokens")
	}
	alllowList := stakingtypes.StakeAuthorization_AllowList{AllowList: &stakingtypes.StakeAuthorization_Validators{Address: info.StakingValidatorAllowList}}

	stakingAuth, err := codectypes.NewAnyWithValue(&stakingtypes.StakeAuthorization{
		MaxTokens:         &sdk.Coin{Denom: config.BaseDenom, Amount: stakingMaxTokens},
		Validators:        &alllowList,
		AuthorizationType: stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
	})
	if err != nil {
		panic(err)
	}
	grants = append(grants, authz.GrantAuthorization{
		Granter:       info.ObserverAddress,
		Grantee:       info.StakingGranteeAddress,
		Authorization: stakingAuth,
		Expiration:    nil,
	})
	delAuth, err := codectypes.NewAnyWithValue(&stakingtypes.StakeAuthorization{
		MaxTokens:         &sdk.Coin{Denom: config.BaseDenom, Amount: stakingMaxTokens},
		Validators:        &alllowList,
		AuthorizationType: stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE,
	})
	if err != nil {
		panic(err)
	}
	grants = append(grants, authz.GrantAuthorization{
		Granter:       info.ObserverAddress,
		Grantee:       info.StakingGranteeAddress,
		Authorization: delAuth,
		Expiration:    nil,
	})
	reDelauth, err := codectypes.NewAnyWithValue(&stakingtypes.StakeAuthorization{
		MaxTokens:         &sdk.Coin{Denom: config.BaseDenom, Amount: stakingMaxTokens},
		Validators:        &alllowList,
		AuthorizationType: stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE,
	})
	if err != nil {
		panic(err)
	}
	grants = append(grants, authz.GrantAuthorization{
		Granter:       info.ObserverAddress,
		Grantee:       info.StakingGranteeAddress,
		Authorization: reDelauth,
		Expiration:    nil,
	})
	return grants

}

// AddObserverAccountCmd Deprecated : Use AddObserverAccountsCmd instead
func AddObserverAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-observer [chainID] [comma separate list of address] ",
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
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			chainID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			chain := common.GetChainFromChainID(chainID)
			observer := &types.ObserverMapper{
				Index:         "",
				ObserverChain: chain,
				ObserverList:  strings.Split(args[1], ","),
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

func AddGenesisAccount(clientCtx client.Context, balances []banktypes.Balance, appState map[string]json.RawMessage) (map[string]json.RawMessage, error) {
	var genAccount authtypes.GenesisAccount
	totalBalanceAdded := sdk.Coins{}
	genAccounts := make([]authtypes.GenesisAccount, len(balances))
	for i, balance := range balances {
		totalBalanceAdded = totalBalanceAdded.Add(balance.Coins...)
		accAddress := sdk.MustAccAddressFromBech32(balance.Address)
		baseAccount := authtypes.NewBaseAccount(accAddress, nil, 0, 0)
		genAccount = &ethermint.EthAccount{
			BaseAccount: baseAccount,
			CodeHash:    ethcommon.BytesToHash(evmtypes.EmptyCodeHash).Hex(),
		}
		if err := genAccount.Validate(); err != nil {
			return appState, fmt.Errorf("failed to validate new genesis account: %w", err)
		}
		genAccounts[i] = genAccount
	}

	authGenState := authtypes.GetGenesisStateFromAppState(clientCtx.Codec, appState)

	accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
	if err != nil {
		return appState, fmt.Errorf("failed to get accounts from any: %w", err)
	}

	for _, genAc := range genAccounts {
		addr := genAc.GetAddress()
		if accs.Contains(addr) {
			return appState, fmt.Errorf("cannot add account at existing address %s", addr)
		}
		accs = append(accs, genAc)
		accs = authtypes.SanitizeGenesisAccounts(accs)
	}

	genAccs, err := authtypes.PackAccounts(accs)
	if err != nil {
		return appState, fmt.Errorf("failed to convert accounts into any's: %w", err)
	}
	authGenState.Accounts = genAccs

	authGenStateBz, err := clientCtx.Codec.MarshalJSON(&authGenState)
	if err != nil {
		return appState, fmt.Errorf("failed to marshal auth genesis state: %w", err)
	}
	appState[authtypes.ModuleName] = authGenStateBz
	bankGenState := banktypes.GetGenesisStateFromAppState(clientCtx.Codec, appState)
	bankGenState.Balances = append(bankGenState.Balances, balances...)
	bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)
	bankGenState.Supply = bankGenState.Supply.Add(totalBalanceAdded...)

	bankGenStateBz, err := clientCtx.Codec.MarshalJSON(bankGenState)
	if err != nil {
		return appState, fmt.Errorf("failed to marshal bank genesis state: %w", err)
	}
	appState[banktypes.ModuleName] = bankGenStateBz

	return appState, nil
}

func isValidatorOnly(isObserver string) bool {
	if isObserver == "Y" {
		return false
	} else if isObserver == "N" {
		return true
	}
	panic("Invalid Input for isObserver field, Check observer_info.json file")
}
