package zetaclient

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	zetaObserverModuleTypes "github.com/zeta-chain/zetacore/x/observer/types"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type EVMSigner struct {
	client                      *ethclient.Client
	chain                       *common.Chain
	chainID                     *big.Int
	tssSigner                   TSSSigner
	ethSigner                   ethtypes.Signer
	abi                         abi.ABI
	erc20CustodyABI             abi.ABI
	metaContractAddress         ethcommon.Address
	erc20CustodyContractAddress ethcommon.Address
	logger                      zerolog.Logger
}

var _ ChainSigner = &EVMSigner{}

func NewEVMSigner(chain common.Chain, endpoint string, tssSigner TSSSigner, abiString string, erc20CustodyABIString string, metaContract ethcommon.Address, erc20CustodyContract ethcommon.Address) (*EVMSigner, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}
	ethSigner := ethtypes.LatestSignerForChainID(chainID)
	connectorABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}
	erc20CustodyABI, err := abi.JSON(strings.NewReader(erc20CustodyABIString))
	if err != nil {
		return nil, err
	}

	return &EVMSigner{
		client:                      client,
		chain:                       &chain,
		tssSigner:                   tssSigner,
		chainID:                     chainID,
		ethSigner:                   ethSigner,
		abi:                         connectorABI,
		erc20CustodyABI:             erc20CustodyABI,
		metaContractAddress:         metaContract,
		erc20CustodyContractAddress: erc20CustodyContract,
		logger:                      log.With().Str("module", "EVMSigner").Logger(),
	}, nil
}

// given data, and metadata (gas, nonce, etc)
// returns a signed transaction, sig bytes, hash bytes, and error
func (signer *EVMSigner) Sign(data []byte, to ethcommon.Address, gasLimit uint64, gasPrice *big.Int, nonce uint64) (*ethtypes.Transaction, []byte, []byte, error) {
	tx := ethtypes.NewTransaction(nonce, to, big.NewInt(0), gasLimit, gasPrice, data)
	hashBytes := signer.ethSigner.Hash(tx).Bytes()
	sig, err := signer.tssSigner.Sign(hashBytes)
	if err != nil {
		return nil, nil, nil, err
	}
	pubk, err := crypto.SigToPub(hashBytes, sig[:])
	if err != nil {
		signer.logger.Error().Err(err).Msgf("SigToPub error")
	}
	addr := crypto.PubkeyToAddress(*pubk)
	signer.logger.Info().Msgf("Sign: Ecrecovery of signature: %s", addr.Hex())
	signedTX, err := tx.WithSignature(signer.ethSigner, sig[:])
	if err != nil {
		return nil, nil, nil, err
	}
	return signedTX, sig[:], hashBytes[:], nil
}

// takes in signed tx, broadcast to external chain node
func (signer *EVMSigner) Broadcast(tx *ethtypes.Transaction) error {
	ctxt, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return signer.client.SendTransaction(ctxt, tx)
}

// function onReceive(
//
//	bytes calldata originSenderAddress,
//	uint256 originChainId,
//	address destinationAddress,
//	uint zetaAmount,
//	bytes calldata message,
//	bytes32 internalSendHash
//
// ) external virtual {}
func (signer *EVMSigner) SignOutboundTx(sender ethcommon.Address, srcChainID *big.Int, to ethcommon.Address, amount *big.Int, gasLimit uint64, message []byte, sendHash [32]byte, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	if len(sendHash) < 32 {
		return nil, fmt.Errorf("sendHash len %d must be 32", len(sendHash))
	}
	var data []byte
	var err error

	data, err = signer.abi.Pack("onReceive", sender.Bytes(), srcChainID, to, amount, message, sendHash)
	if err != nil {
		return nil, fmt.Errorf("pack error: %w", err)
	}

	tx, _, _, err := signer.Sign(data, signer.metaContractAddress, gasLimit, gasPrice, nonce)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %w", err)
	}

	return tx, nil
}

// function onRevert(
// address originSenderAddress,
// uint256 originChainId,
// bytes calldata destinationAddress,
// uint256 destinationChainId,
// uint256 zetaAmount,
// bytes calldata message,
// bytes32 internalSendHash
// ) external override whenNotPaused onlyTssAddress
func (signer *EVMSigner) SignRevertTx(sender ethcommon.Address, srcChainID *big.Int, to []byte, toChainID *big.Int, amount *big.Int, gasLimit uint64, message []byte, sendHash [32]byte, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	var data []byte
	var err error

	data, err = signer.abi.Pack("onRevert", sender, srcChainID, to, toChainID, amount, message, sendHash)
	if err != nil {
		return nil, fmt.Errorf("pack error: %w", err)
	}

	tx, _, _, err := signer.Sign(data, signer.metaContractAddress, gasLimit, gasPrice, nonce)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %w", err)
	}

	return tx, nil
}

func (signer *EVMSigner) SignCancelTx(nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	tx := ethtypes.NewTransaction(nonce, signer.tssSigner.EVMAddress(), big.NewInt(0), 21000, gasPrice, nil)
	hashBytes := signer.ethSigner.Hash(tx).Bytes()
	sig, err := signer.tssSigner.Sign(hashBytes)
	if err != nil {
		return nil, err
	}
	pubk, err := crypto.SigToPub(hashBytes, sig[:])
	if err != nil {
		signer.logger.Error().Err(err).Msgf("SigToPub error")
	}
	addr := crypto.PubkeyToAddress(*pubk)
	signer.logger.Info().Msgf("Sign: Ecrecovery of signature: %s", addr.Hex())
	signedTX, err := tx.WithSignature(signer.ethSigner, sig[:])
	if err != nil {
		return nil, err
	}

	return signedTX, nil
}

func (signer *EVMSigner) SignWithdrawTx(to ethcommon.Address, amount *big.Int, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	tx := ethtypes.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)
	hashBytes := signer.ethSigner.Hash(tx).Bytes()
	sig, err := signer.tssSigner.Sign(hashBytes)
	if err != nil {
		return nil, err
	}
	pubk, err := crypto.SigToPub(hashBytes, sig[:])
	if err != nil {
		signer.logger.Error().Err(err).Msgf("SigToPub error")
	}
	addr := crypto.PubkeyToAddress(*pubk)
	signer.logger.Info().Msgf("Sign: Ecrecovery of signature: %s", addr.Hex())
	signedTX, err := tx.WithSignature(signer.ethSigner, sig[:])
	if err != nil {
		return nil, err
	}

	return signedTX, nil
}

func (signer *EVMSigner) TryProcessOutTx(send *types.CrossChainTx, outTxMan *OutTxProcessorManager, outTxID string, evmClient ChainClient, zetaBridge *ZetaCoreBridge) {
	logger := signer.logger.With().
		Str("sendHash", send.Index).
		Str("outTxID", outTxID).
		Logger()
	logger.Info().Msgf("start processing outTxID %s", outTxID)
	defer func() {
		outTxMan.EndTryProcess(outTxID)
	}()
	myid := zetaBridge.keys.GetAddress().String()

	var to ethcommon.Address
	var err error
	var toChain *common.Chain
	if send.CctxStatus.Status == types.CctxStatus_PendingRevert {
		to = ethcommon.HexToAddress(send.InboundTxParams.Sender)
		toChain = GetChainFromChainID(send.InboundTxParams.SenderChainId)
		logger.Info().Msgf("Abort: reverting inbound")
	} else if send.CctxStatus.Status == types.CctxStatus_PendingOutbound {
		to = ethcommon.HexToAddress(send.GetCurrentOutTxParam().Receiver)
		toChain = GetChainFromChainID(send.GetCurrentOutTxParam().ReceiverChainId)
	}
	if err != nil {
		logger.Error().Err(err).Msg("ParseChain fail; skip")
		return
	}

	// Early return if the send is already processed
	included, confirmed, _ := evmClient.IsSendOutTxProcessed(send.Index, int(send.GetCurrentOutTxParam().OutboundTxTssNonce), send.GetCurrentOutTxParam().CoinType)
	if included || confirmed {
		logger.Info().Msgf("CCTX already processed; exit signer")
		return
	}

	message, err := base64.StdEncoding.DecodeString(send.RelayedMessage)
	if err != nil {
		logger.Err(err).Msgf("decode CCTX.Message %s error", send.RelayedMessage)
	}

	gasLimit := send.GetCurrentOutTxParam().OutboundTxGasLimit
	if gasLimit < 50_000 {
		gasLimit = 50_000
		logger.Warn().Msgf("gasLimit %d is too low; set to %d", send.GetCurrentOutTxParam().OutboundTxGasLimit, gasLimit)
	}
	if gasLimit > 1_000_000 {
		gasLimit = 1_000_000
		logger.Warn().Msgf("gasLimit %d is too high; set to %d", send.GetCurrentOutTxParam().OutboundTxGasLimit, gasLimit)
	}

	logger.Info().Msgf("chain %s minting %d to %s, nonce %d, finalized zeta bn %d", toChain, send.InboundTxParams.Amount, to.Hex(), send.GetCurrentOutTxParam().OutboundTxTssNonce, send.InboundTxParams.InboundTxFinalizedZetaHeight)
	sendHash, err := hex.DecodeString(send.Index[2:]) // remove the leading 0x
	if err != nil || len(sendHash) != 32 {
		logger.Error().Err(err).Msgf("decode CCTX %s error", send.Index)
		return
	}
	var sendhash [32]byte
	copy(sendhash[:32], sendHash[:32])
	gasprice, ok := new(big.Int).SetString(send.GetCurrentOutTxParam().OutboundTxGasPrice, 10)
	if !ok {
		logger.Error().Err(err).Msgf("cannot convert gas price  %s ", send.GetCurrentOutTxParam().OutboundTxGasPrice)
		return
	}

	var tx *ethtypes.Transaction
	// FIXME: there is a chance wrong type of outbound tx is signed
	// NOTE: if sender is zetachain, then the tx could be one of three types;
	// otherwise, it's always a message passing tx, passing zeta & optionally message
	if send.InboundTxParams.SenderChainId == common.ZetaChain().ChainId && send.CctxStatus.Status == types.CctxStatus_PendingOutbound {
		if send.GetCurrentOutTxParam().CoinType == common.CoinType_Gas {
			logger.Info().Msgf("SignWithdrawTx: %d => %s, nonce %d, gasprice %d", send.InboundTxParams.SenderChainId, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
			tx, err = signer.SignWithdrawTx(to, send.InboundTxParams.Amount.BigInt(), send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
		}
		if send.GetCurrentOutTxParam().CoinType == common.CoinType_ERC20 {
			asset := ethcommon.HexToAddress(send.InboundTxParams.Asset)
			logger.Info().Msgf("SignERC20WithdrawTx: %d => %s, nonce %d, gasprice %d", send.InboundTxParams.SenderChainId, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
			tx, err = signer.SignERC20WithdrawTx(to, asset, send.InboundTxParams.Amount.BigInt(), gasLimit, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
		}
		if send.GetCurrentOutTxParam().CoinType == common.CoinType_Zeta {
			//srcChainID := config.ChainConfigs[send.InboundTxParams.SenderChainId].Chain.ChainId
			logger.Info().Msgf("SignOutboundTx: %d => %s, nonce %d, gasprice %d", send.InboundTxParams.SenderChainId, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
			tx, err = signer.SignOutboundTx(ethcommon.HexToAddress(send.InboundTxParams.Sender), big.NewInt(send.InboundTxParams.SenderChainId), to, send.InboundTxParams.Amount.BigInt(), gasLimit, message, sendhash, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
		}
	} else if send.CctxStatus.Status == types.CctxStatus_PendingRevert {
		//srcChainID := config.ChainConfigs[send.InboundTxParams.SenderChain].Chain.ChainId
		logger.Info().Msgf("SignRevertTx: %d => %s, nonce %d, gasprice %d", send.InboundTxParams.SenderChainId, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
		//toChainID := config.ChainConfigs[send.OutboundTxParams.ReceiverChain].Chain.ChainId
		tx, err = signer.SignRevertTx(ethcommon.HexToAddress(send.InboundTxParams.Sender), big.NewInt(send.OutboundTxParams[0].ReceiverChainId), to.Bytes(), big.NewInt(send.GetCurrentOutTxParam().ReceiverChainId), send.GetCurrentOutTxParam().Amount.BigInt(), gasLimit, message, sendhash, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
	} else if send.CctxStatus.Status == types.CctxStatus_PendingOutbound {
		logger.Info().Msgf("SignOutboundTx: %d => %s, nonce %d, gasprice %d", send.InboundTxParams.SenderChainId, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
		tx, err = signer.SignOutboundTx(ethcommon.HexToAddress(send.InboundTxParams.Sender), big.NewInt(send.InboundTxParams.SenderChainId), to, send.GetCurrentOutTxParam().Amount.BigInt(), gasLimit, message, sendhash, send.GetCurrentOutTxParam().OutboundTxTssNonce, gasprice)
	}

	if err != nil {
		logger.Warn().Err(err).Msgf("signer SignOutbound error: nonce %d chain %d", send.GetCurrentOutTxParam().OutboundTxTssNonce, send.GetCurrentOutTxParam().ReceiverChainId)
		return
	}
	logger.Info().Msgf("Key-sign success: %d => %s, nonce %d", send.InboundTxParams.SenderChainId, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce)

	_, err = zetaBridge.GetObserverList(*toChain, zetaObserverModuleTypes.ObservationType_OutBoundTx.String())
	if err != nil {
		logger.Warn().Err(err).Msgf("unable to get observer list: chain %d observation %s", send.GetCurrentOutTxParam().OutboundTxTssNonce, zetaObserverModuleTypes.ObservationType_OutBoundTx.String())

	}
	if tx != nil {
		outTxHash := tx.Hash().Hex()
		logger.Info().Msgf("on chain %s nonce %d, outTxHash %s signer %s", signer.chain, send.GetCurrentOutTxParam().OutboundTxTssNonce, outTxHash, myid)
		//if len(signers) == 0 || myid == signers[send.OutboundTxParams.Broadcaster] || myid == signers[int(send.OutboundTxParams.Broadcaster+1)%len(signers)] {
		backOff := 1000 * time.Millisecond
		// retry loop: 1s, 2s, 4s, 8s, 16s in case of RPC error
		for i := 0; i < 5; i++ {
			logger.Info().Msgf("broadcasting tx %s to chain %s: nonce %d, retry %d", outTxHash, toChain, send.GetCurrentOutTxParam().OutboundTxTssNonce, i)
			// #nosec G404 randomness is not a security issue here
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // FIXME: use backoff
			err := signer.Broadcast(tx)
			if err != nil {
				log.Warn().Err(err).Msgf("OutTx Broadcast error")
				retry, report := HandleBroadcastError(err, strconv.FormatUint(send.GetCurrentOutTxParam().OutboundTxTssNonce, 10), toChain.String(), outTxHash)
				if report {
					zetaHash, err := zetaBridge.AddTxHashToOutTxTracker(toChain.ChainId, tx.Nonce(), outTxHash)
					if err != nil {
						logger.Err(err).Msgf("Unable to add to tracker on ZetaCore: nonce %d chain %s outTxHash %s", send.GetCurrentOutTxParam().OutboundTxTssNonce, toChain, outTxHash)
					}
					logger.Info().Msgf("Broadcast to core successful %s", zetaHash)
				}
				if !retry {
					break
				}
				backOff *= 2
				continue
			}
			logger.Info().Msgf("Broadcast success: nonce %d to chain %s outTxHash %s", send.GetCurrentOutTxParam().OutboundTxTssNonce, toChain, outTxHash)
			zetaHash, err := zetaBridge.AddTxHashToOutTxTracker(toChain.ChainId, tx.Nonce(), outTxHash)
			if err != nil {
				logger.Err(err).Msgf("Unable to add to tracker on ZetaCore: nonce %d chain %s outTxHash %s", send.GetCurrentOutTxParam().OutboundTxTssNonce, toChain, outTxHash)
			}
			logger.Info().Msgf("Broadcast to core successful %s", zetaHash)
			break // successful broadcast; no need to retry
		}

	}
	//}

}

// function withdraw(
// address recipient,
// address asset,
// uint256 amount,
// ) external onlyTssAddress
func (signer *EVMSigner) SignERC20WithdrawTx(recipient ethcommon.Address, asset ethcommon.Address, amount *big.Int, gasLimit uint64, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	var data []byte
	var err error
	data, err = signer.erc20CustodyABI.Pack("withdraw", recipient, asset, amount)
	if err != nil {
		return nil, fmt.Errorf("pack error: %w", err)
	}

	tx, _, _, err := signer.Sign(data, signer.erc20CustodyContractAddress, gasLimit, gasPrice, nonce)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %w", err)
	}

	return tx, nil
}

// function whitelist(
// address asset,
// ) external onlyTssAddress
// function unwhitelist(
// address asset,
// ) external onlyTssAddress
func (signer *EVMSigner) SignWhitelistTx(action string, recipient ethcommon.Address, asset ethcommon.Address, gasLimit uint64, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {
	var data []byte

	var err error

	data, err = signer.erc20CustodyABI.Pack(action, asset)
	if err != nil {
		return nil, fmt.Errorf("pack error: %w", err)
	}

	tx, _, _, err := signer.Sign(data, signer.erc20CustodyContractAddress, gasLimit, gasPrice, nonce)
	if err != nil {
		return nil, fmt.Errorf("Sign error: %w", err)
	}

	return tx, nil
}
