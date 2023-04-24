package main

import (
	"fmt"
	"github.com/btcsuite/btcd/rpcclient"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/common/cosmos"
	"github.com/zeta-chain/zetacore/zetaclient"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	"github.com/zeta-chain/zetacore/zetaclient/metrics"
)

func CreateAuthzSigner(granter string, grantee sdk.AccAddress) {
	zetaclient.SetupAuthZSignerList(granter, grantee)
}

func CreateZetaBridge(chainHomeFolder string, config *config.Config) (*zetaclient.ZetaCoreBridge, error) {
	signerPass := "password"
	chainIP := config.ZetaCoreURL
	kb, err := zetaclient.GetKeyringKeybase(config.AuthzHotkey, chainHomeFolder, signerPass)
	if err != nil {
		return nil, err
	}
	granterAddreess, err := cosmos.AccAddressFromBech32(config.AuthzGranter)
	if err != nil {
		return nil, err
	}
	k := zetaclient.NewKeysWithKeybase(kb, granterAddreess, config.AuthzHotkey, signerPass)
	bridge, err := zetaclient.NewZetaCoreBridge(k, chainIP, config.AuthzHotkey)
	if err != nil {
		return nil, err
	}
	return bridge, nil
}

func CreateSignerMap(tss zetaclient.TSSSigner, logger zerolog.Logger, cfg *config.Config) (map[common.Chain]zetaclient.ChainSigner, error) {
	signerMap := make(map[common.Chain]zetaclient.ChainSigner)
	for _, chain := range cfg.ChainsEnabled {
		if chain.IsZetaChain() {
			continue
		}
		if chain.IsEVMChain() {
			mpiAddress := ethcommon.HexToAddress(cfg.EVMChainConfigs[chain.ChainName.String()].CoreParams.ConnectorContractAddress)
			erc20CustodyAddress := ethcommon.HexToAddress(cfg.EVMChainConfigs[chain.ChainName.String()].CoreParams.ERC20CustodyContractAddress)
			signer, err := zetaclient.NewEVMSigner(chain, cfg.EVMChainConfigs[chain.ChainName.String()].Endpoint, tss, config.GetConnectorABI(), config.GetERC20CustodyABI(), mpiAddress, erc20CustodyAddress, logger)
			if err != nil {
				logger.Fatal().Err(err).Msgf("%s: NewEVMSigner Ethereum error ", chain.String())
				return nil, err
			}
			signerMap[chain] = signer
		} else if chain.IsBitcoinChain() {
			// FIXME: move the construction of rpcclient to somewhere else
			connCfg := &rpcclient.ConnConfig{
				Host:         cfg.BitcoinConfig.RPCEndpoint,
				User:         cfg.BitcoinConfig.RPCUsername,
				Pass:         cfg.BitcoinConfig.RPCPassword,
				HTTPPostMode: true,
				DisableTLS:   true,
				Params:       cfg.BitcoinConfig.RPCParams,
			}
			client, err := rpcclient.New(connCfg, nil)
			if err != nil {
				return nil, fmt.Errorf("error creating rpc client: %s", err)
			}
			signer, err := zetaclient.NewBTCSigner(tss, client, logger)
			if err != nil {
				logger.Fatal().Err(err).Msgf("%s: NewBitcoinSigner Bitcoin error ", chain.String())
				return nil, err
			}
			signerMap[chain] = signer
		}
	}

	return signerMap, nil
}

func CreateChainClientMap(bridge *zetaclient.ZetaCoreBridge, tss zetaclient.TSSSigner, dbpath string, metrics *metrics.Metrics, logger zerolog.Logger, cfg *config.Config) (map[common.Chain]zetaclient.ChainClient, error) {
	clientMap := make(map[common.Chain]zetaclient.ChainClient)
	for _, chain := range cfg.ChainsEnabled {
		if chain.IsZetaChain() {
			continue
		}
		logger.Info().Msgf("starting observer for : %s ", chain.String())
		var co zetaclient.ChainClient
		var err error
		if chain.IsEVMChain() {
			co, err = zetaclient.NewEVMChainClient(bridge, tss, dbpath, metrics, logger, cfg, chain)
		} else if chain.IsBitcoinChain() {
			co, err = zetaclient.NewBitcoinClient(chain, bridge, tss, dbpath, metrics, logger)
		}
		if err != nil {
			log.Err(err).Msgf("%s NewEVMChainClient", chain.String())
			return nil, err
		}
		clientMap[chain] = co
	}

	return clientMap, nil
}
