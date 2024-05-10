// TODO: use with signer extractor once available https://github.com/zeta-chain/node/issues/2156

package mempool

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// GetSendersWithNonce is used to extract sender and nonce information txs
// if tx is ethermint, it is extracted using from and nonce field
// if it's cosmos tx, default cosmos way using signatures is used
func GetSendersWithNonce(tx sdk.Tx) ([]SenderWithNonce, error) {
	const extensionOptionsEthereumTxTypeUrl = "/ethermint.evm.v1.ExtensionOptionsEthereumTx"
	if txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx); ok {
		opts := txWithExtensions.GetExtensionOptions()
		if len(opts) > 0 && opts[0].GetTypeUrl() == extensionOptionsEthereumTxTypeUrl {
			for _, msg := range tx.GetMsgs() {
				if ethMsg, ok := msg.(*evmtypes.MsgEthereumTx); ok {
					return []SenderWithNonce{
						{
							Sender: ethMsg.GetFrom().String(),
							Nonce:  ethMsg.AsTransaction().Nonce(),
						},
					}, nil
				}
			}
		}
	}

	return getSendersWithNonceDefault(tx)
}

type SenderWithNonce struct {
	Sender string
	Nonce  uint64
}

// getSendersWithNonceDefault gets senders and nonces from signatures in cosmos txs
func getSendersWithNonceDefault(tx sdk.Tx) ([]SenderWithNonce, error) {
	sendersWithNonce := []SenderWithNonce{}

	sigTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return nil, fmt.Errorf("tx of type %T does not implement SigVerifiableTx", tx)
	}

	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return nil, err
	}

	if len(sigs) == 0 {
		return nil, fmt.Errorf("tx must have at least one signer")
	}

	for _, sig := range sigs {
		sendersWithNonce = append(sendersWithNonce, SenderWithNonce{
			Sender: sig.PubKey.Address().String(),
			Nonce:  sig.Sequence,
		})
	}

	return sendersWithNonce, nil
}
