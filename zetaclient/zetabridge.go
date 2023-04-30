package zetaclient

import (
	"fmt"
	"github.com/zeta-chain/zetacore/common"
	"time"

	"sync"

	"github.com/hashicorp/go-retryablehttp"
	"google.golang.org/grpc"

	//"fmt"
	"github.com/zeta-chain/zetacore/common/cosmos"
	//"github.com/armon/go-metrics"
	//"github.com/cosmos/cosmos-sdk/Client"
	"github.com/cosmos/cosmos-sdk/codec"

	//"github.com/cosmos/cosmos-sdk/std"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	//"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	//"golang.org/x/tools/go/cfg"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"strconv"
	//"strings"

	stypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	"github.com/zeta-chain/zetacore/zetaclient/config"
)

// ZetaCoreBridge will be used to send tx to ZetaCore.
type ZetaCoreBridge struct {
	logger        zerolog.Logger
	blockHeight   int64
	accountNumber map[common.KeyType]uint64
	seqNumber     map[common.KeyType]uint64
	grpcConn      *grpc.ClientConn
	httpClient    *retryablehttp.Client
	cfg           config.ClientConfiguration
	keys          *Keys
	broadcastLock *sync.RWMutex
	//ChainNonces         map[string]uint64 // FIXME: Remove this?
	lastOutTxReportTime map[string]time.Time
	stop                chan struct{}
}

// NewZetaCoreBridge create a new instance of ZetaCoreBridge
func NewZetaCoreBridge(k *Keys, chainIP string, signerName string) (*ZetaCoreBridge, error) {
	// main module logger
	logger := log.With().Str("module", "CoreBridge").Logger()
	cfg := config.ClientConfiguration{
		ChainHost:    fmt.Sprintf("%s:1317", chainIP),
		SignerName:   signerName,
		SignerPasswd: "password",
		ChainRPC:     fmt.Sprintf("%s:26657", chainIP),
	}

	httpClient := retryablehttp.NewClient()
	httpClient.Logger = nil

	grpcConn, err := grpc.Dial(
		fmt.Sprintf("%s:9090", chainIP),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Error().Err(err).Msg("grpc dial fail")
		return nil, err
	}
	accountsMap := make(map[common.KeyType]uint64)
	seqMap := make(map[common.KeyType]uint64)
	for _, keyType := range common.GetAllKeyTypes() {
		accountsMap[keyType] = 0
		seqMap[keyType] = 0
	}

	return &ZetaCoreBridge{
		logger:              logger,
		grpcConn:            grpcConn,
		httpClient:          httpClient,
		accountNumber:       accountsMap,
		seqNumber:           seqMap,
		cfg:                 cfg,
		keys:                k,
		broadcastLock:       &sync.RWMutex{},
		lastOutTxReportTime: map[string]time.Time{},
		stop:                make(chan struct{}),
	}, nil
}

// MakeLegacyCodec creates codec
func MakeLegacyCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	banktypes.RegisterLegacyAminoCodec(cdc)
	authtypes.RegisterLegacyAminoCodec(cdc)
	cosmos.RegisterCodec(cdc)
	stypes.RegisterCodec(cdc)
	return cdc
}

func (b *ZetaCoreBridge) Stop() {
	b.logger.Info().Msgf("ZetaBridge is stopping")
	close(b.stop) // this notifies all configupdater to stop
}

func (b *ZetaCoreBridge) GetAccountNumberAndSequenceNumber(keyType common.KeyType) (uint64, uint64, error) {
	ctx := b.GetContext()
	address := b.keys.GetAddress()
	return ctx.AccountRetriever.GetAccountNumberSequence(ctx, address)
}

func (b *ZetaCoreBridge) SetAccountNumber(keyType common.KeyType) {
	ctx := b.GetContext()
	address := b.keys.GetAddress()
	accN, seq, _ := ctx.AccountRetriever.GetAccountNumberSequence(ctx, address)
	b.accountNumber[keyType] = accN
	b.seqNumber[keyType] = seq
}

func (b *ZetaCoreBridge) WaitForCoreToCreateBlocks() {
	retryCount := 0
	maxRetryCount := 10
	for {
		block, err := b.GetLatestZetaBlock()
		if err == nil && block.Header.Height > 1 {
			b.logger.Info().Msgf("Zeta-core height: %d", block.Header.Height)
			break
		}
		retryCount++
		b.logger.Debug().Msgf("Failed to get latest Block , Retry : %d/%d", retryCount, maxRetryCount)
		if retryCount > maxRetryCount {
			panic("ZetaCore is not ready , Waited for 60s")
		}
		time.Sleep(6 * time.Second)
	}
}

//func (b *ZetaCoreBridge) GetOperatorAccountNumberAndSequenceNumber() (uint64, uint64, error) {
//	ctx := b.GetContext()
//	return ctx.AccountRetriever.GetAccountNumberSequence(ctx, b.keys.GetOperatorAddress())
//}

func (b *ZetaCoreBridge) GetKeys() *Keys {
	return b.keys
}

func (b *ZetaCoreBridge) UpdateConfigFromCore(config *config.Config) error {
	coreParams, err := b.GetCoreParams()
	if err != nil {
		return err
	}
	chains := make([]common.Chain, len(coreParams))
	for i, params := range coreParams {
		chains[i] = *common.GetChainFromChainID(params.ChainId)
		if common.IsBitcoinChain(params.ChainId) {
			config.BitcoinConfig.CoreParams.UpdateCoreParams(params)
			continue
		}
		config.EVMChainConfigs[params.ChainId].CoreParams.UpdateCoreParams(params)
	}
	config.ChainsEnabled = chains

	return nil
}
