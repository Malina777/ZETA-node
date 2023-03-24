package zetaclient

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"math/big"
	"time"

	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

const (
	PostGasPriceGasLimit            = 1_500_000
	AddTxHashToOutTxTrackerGasLimit = 200_000
	PostNonceGasLimit               = 200_000
	PostSendEVMGasLimit             = 1_000_000 // likely emit a lot of logs, so costly
	PostSendNonEVMGasLimit          = 1_000_000
	PostReceiveConfirmationGasLimit = 200_000
	DefaultGasLimit                 = 200_000
)

func (b *ZetaCoreBridge) WrapMessageWithAuthz(msg sdk.Msg) (sdk.Msg, AuthZSigner) {
	msgURL := sdk.MsgTypeURL(msg)
	authzSigner := GetSigner(msgURL)
	authzMessage := authz.NewMsgExec(authzSigner.GranteeAddress, []sdk.Msg{msg})
	return &authzMessage, authzSigner
}

func (b *ZetaCoreBridge) PostGasPrice(chain common.Chain, gasPrice uint64, supply string, blockNum uint64) (string, error) {

	signerAddress := b.keys.GetOperatorAddress().String()
	msg := types.NewMsgGasPriceVoter(signerAddress, chain.ChainId, gasPrice, supply, blockNum)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	zetaTxHash, err := b.Broadcast(PostGasPriceGasLimit, authzMsg, authzSigner)
	if err != nil {
		return "", err
	}

	return zetaTxHash, nil
}

func (b *ZetaCoreBridge) AddTxHashToOutTxTracker(chainID int64, nonce uint64, txHash string) (string, error) {
	signerAddress := b.keys.GetOperatorAddress().String()
	msg := types.NewMsgAddToOutTxTracker(signerAddress, chainID, nonce, txHash)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	zetaTxHash, err := b.Broadcast(AddTxHashToOutTxTrackerGasLimit, authzMsg, authzSigner)
	if err != nil {
		return "", err
	}
	return zetaTxHash, nil
}

func (b *ZetaCoreBridge) PostNonce(chain common.Chain, nonce uint64) (string, error) {
	signerAddress := b.keys.GetOperatorAddress().String()
	msg := types.NewMsgNonceVoter(signerAddress, chain.ChainId, nonce)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	zetaTxHash, err := b.Broadcast(PostNonceGasLimit, authzMsg, authzSigner)
	if err != nil {
		return "", err
	}
	return zetaTxHash, nil
}

func (b *ZetaCoreBridge) PostSend(sender string, senderChain int64, txOrigin string, receiver string, receiverChain int64, amount math.Uint, message string, inTxHash string, inBlockHeight uint64, gasLimit uint64, coinType common.CoinType, zetaGasLimit uint64, asset string) (string, error) {
	signerAddress := b.keys.GetOperatorAddress().String()
	msg := types.NewMsgSendVoter(signerAddress, sender, senderChain, txOrigin, receiver, receiverChain, amount, message, inTxHash, inBlockHeight, gasLimit, coinType, asset)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	maxRetries := 2
	for i := 0; i < maxRetries; i++ {
		zetaTxHash, err := b.Broadcast(zetaGasLimit, authzMsg, authzSigner)
		if err == nil {
			return zetaTxHash, nil
		}
		b.logger.Debug().Err(err).Msgf("PostSend broadcast fail | Retry count : %d", i+1)
		time.Sleep(1 * time.Second)
	}

	return "", fmt.Errorf("post send failed after %d retries", maxRetries)
}

func (b *ZetaCoreBridge) PostReceiveConfirmation(sendHash string, outTxHash string, outBlockHeight uint64, amount *big.Int, status common.ReceiveStatus, chain common.Chain, nonce int, coinType common.CoinType) (string, error) {
	lastReport, found := b.lastOutTxReportTime[outTxHash]
	if found && time.Since(lastReport) < 10*time.Minute {
		return "", fmt.Errorf("PostReceiveConfirmation: outTxHash %s already reported in last 10min; last report %s", outTxHash, lastReport)
	}

	signerAddress := b.keys.GetOperatorAddress().String()
	msg := types.NewMsgReceiveConfirmation(signerAddress, sendHash, outTxHash, outBlockHeight, math.NewUintFromBigInt(amount), status, chain.ChainId, uint64(nonce), coinType)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	// FIXME: remove this gas limit stuff; in the special ante handler with no gas limit, add
	// NewMsgReceiveConfirmation to it.
	var gasLimit uint64 = PostReceiveConfirmationGasLimit
	if status == common.ReceiveStatus_Failed {
		gasLimit = PostSendEVMGasLimit
	}
	maxRetries := 2
	for i := 0; i < maxRetries; i++ {
		zetaTxHash, err := b.Broadcast(gasLimit, authzMsg, authzSigner)
		if err == nil {
			b.lastOutTxReportTime[outTxHash] = time.Now() // update last report time when bcast succeeds
			return zetaTxHash, nil
		}
		b.logger.Debug().Err(err).Msgf("PostReceive broadcast fail | Retry count : %d", i+1)
		time.Sleep(1 * time.Second)
	}
	return "", fmt.Errorf("post receive failed after %d retries", maxRetries)
}

func (b *ZetaCoreBridge) SetNodeKey(tssPubkeySet common.PubKeySet, conskey string) (string, error) {
	operatorAddress := b.keys.GetOperatorAddress()
	msg := types.NewMsgSetNodeKeys(operatorAddress.String(), tssPubkeySet, conskey)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	zetaTxHash, err := b.Broadcast(DefaultGasLimit, authzMsg, authzSigner)
	if err != nil {
		return "", err
	}
	return zetaTxHash, nil
}

func (b *ZetaCoreBridge) SetTSS(chain common.Chain, tssAddress string, tssPubkey string) (string, error) {
	signerAddress := b.keys.GetOperatorAddress().String()
	msg := types.NewMsgCreateTSSVoter(signerAddress, chain.ChainName.String(), tssAddress, tssPubkey)
	authzMsg, authzSigner := b.WrapMessageWithAuthz(msg)
	zetaTxHash, err := b.Broadcast(DefaultGasLimit, authzMsg, authzSigner)
	if err != nil {
		return "", err
	}
	return zetaTxHash, nil
}
