package main

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/common"
	mc "github.com/zeta-chain/zetacore/zetaclient"
	mcconfig "github.com/zeta-chain/zetacore/zetaclient/config"
	"github.com/zeta-chain/zetacore/zetaclient/metrics"
)

func CreateZetaBridge(chainHomeFoler string, signerName string, signerPass string, chainIP string) (*mc.ZetaCoreBridge, bool) {
	kb, _, err := mc.GetKeyringKeybase(chainHomeFoler, signerName, signerPass)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to get keyring keybase")
		return nil, true
	}

	k := mc.NewKeysWithKeybase(kb, signerName, signerPass)

	bridge, err := mc.NewZetaCoreBridge(k, chainIP, signerName)
	if err != nil {
		log.Fatal().Err(err).Msg("NewZetaCoreBridge")
		return nil, true
	}
	return bridge, false
}

func CreateSignerMap(tss mc.TSSSigner) (map[common.Chain]*mc.EVMSigner, error) {
	signerMap := make(map[common.Chain]*mc.EVMSigner)
	supportedChains := common.DefaultChainsList()
	for _, chain := range supportedChains {

		if !(*chain).IsEVMChain() {
			log.Warn().Msgf("chain %s is not an EVM chain, skip creating EVMSigner", chain)
			continue
		}
		mpiAddress := ethcommon.HexToAddress(mcconfig.ChainConfigs[chain.ChainName.String()].ConnectorContractAddress)
		signer, err := mc.NewEVMSigner(chain, mcconfig.ChainConfigs[chain.ChainName.String()].Endpoint, tss, mcconfig.ConnectorAbiString, mpiAddress)
		if err != nil {
			log.Fatal().Err(err).Msgf("%s: NewEVMSigner Ethereum error ", chain.String())
			return nil, err
		}
		signerMap[*chain] = signer
	}

	return signerMap, nil
}

func CreateChainClientMap(bridge *mc.ZetaCoreBridge, tss mc.TSSSigner, dbpath string, metrics *metrics.Metrics) (map[common.Chain]mc.ChainClient, error) {
	clientMap := make(map[common.Chain]mc.ChainClient)
	supportedChains := mc.GetSupportedChains()
	for _, chain := range supportedChains {
		log.Info().Msgf("starting %s observer...", chain)
		var co mc.ChainClient
		var err error
		if chain.IsEVMChain() {
			co, err = mc.NewEVMChainClient(*chain, bridge, tss, dbpath, metrics)
		} else {
			co, err = mc.NewBitcoinClient(*chain, bridge, tss, dbpath, metrics)
		}
		if err != nil {
			log.Err(err).Msgf("%s NewEVMChainClient", chain)
			return nil, err
		}
		clientMap[*chain] = co
	}

	return clientMap, nil
}
