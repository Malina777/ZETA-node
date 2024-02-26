package e2etests

import (
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
)

func TestBitcoinWithdraw(r *runner.E2ERunner) {
	// withdraw 0.1 BTC from ZRC20 to BTC address
	// first, approve the ZRC20 contract to spend 1 BTC from the deployer address
	WithdrawBitcoin(r)
}

func WithdrawBitcoin(r *runner.E2ERunner) {
	amount := big.NewInt(0.1 * btcutil.SatoshiPerBitcoin)

	// approve the ZRC20 contract to spend 1 BTC from the deployer address
	tx, err := r.BTCZRC20.Approve(r.ZevmAuth, r.BTCZRC20Addr, big.NewInt(amount.Int64()*2)) // approve more to cover withdraw fee
	if err != nil {
		panic(err)
	}
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZevmClient, tx, r.Logger, r.ReceiptTimeout)
	if receipt.Status != 1 {
		panic(fmt.Errorf("approve receipt status is not 1"))
	}

	// mine blocks
	stop := r.MineBlocks()

	// withdraw 0.1 BTC from ZRC20 to BTC legacy address
	tx, err = r.BTCZRC20.Withdraw(r.ZevmAuth, []byte("1EYVvXLusCxtVuEwoYvWRyN5EZTXwPVvo3"), amount)
	if err != nil {
		panic(err)
	}
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZevmClient, tx, r.Logger, r.ReceiptTimeout)
	if receipt.Status == 1 {
		panic(fmt.Errorf("withdraw receipt status is successful for an invalid BTC address"))
	}

	// withdraw 0.1 BTC from ZRC20 to BTC address
	tx, err = r.BTCZRC20.Withdraw(r.ZevmAuth, []byte(r.BTCDeployerAddress.EncodeAddress()), amount)
	if err != nil {
		panic(err)
	}
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.ZevmClient, tx, r.Logger, r.ReceiptTimeout)
	if receipt.Status != 1 {
		panic(fmt.Errorf("withdraw receipt status is not 1"))
	}

	// mine 10 blocks to confirm the withdraw tx
	_, err = r.BtcRPCClient.GenerateToAddress(10, r.BTCDeployerAddress, nil)
	if err != nil {
		panic(err)
	}

	cctx := utils.WaitCctxMinedByInTxHash(r.Ctx, receipt.TxHash.Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	outTxHash := cctx.GetCurrentOutTxParam().OutboundTxHash
	hash, err := chainhash.NewHashFromStr(outTxHash)
	if err != nil {
		panic(err)
	}

	rawTx, err := r.BtcRPCClient.GetRawTransactionVerbose(hash)
	if err != nil {
		panic(err)
	}
	r.Logger.Info("raw tx:")
	r.Logger.Info("  TxIn: %d", len(rawTx.Vin))
	for idx, txIn := range rawTx.Vin {
		r.Logger.Info("  TxIn %d:", idx)
		r.Logger.Info("    TxID:Vout:  %s:%d", txIn.Txid, txIn.Vout)
		r.Logger.Info("    ScriptSig: %s", txIn.ScriptSig.Hex)
	}
	r.Logger.Info("  TxOut: %d", len(rawTx.Vout))
	for _, txOut := range rawTx.Vout {
		r.Logger.Info("  TxOut %d:", txOut.N)
		r.Logger.Info("    Value: %.8f", txOut.Value)
		r.Logger.Info("    ScriptPubKey: %s", txOut.ScriptPubKey.Hex)
	}

	// stop mining
	stop <- struct{}{}
}

// WithdrawBitcoinMultipleTimes ...
// TODO: complete and uncomment E2E test
// https://github.com/zeta-chain/node-private/issues/79
//func WithdrawBitcoinMultipleTimes(r *runner.E2ERunner, repeat int64) {
//	totalAmount := big.NewInt(int64(0.1 * 1e8))
//
//	// #nosec G701 test - always in range
//	amount := big.NewInt(int64(0.1 * 1e8 / float64(repeat)))
//
//	// check if the deposit is successful
//	BTCZRC20Addr, err := r.SystemContract.GasCoinZRC20ByChainId(&bind.CallOpts{}, big.NewInt(common.BtcRegtestChain().ChainId))
//	if err != nil {
//		panic(err)
//	}
//	r.Logger.Info("BTCZRC20Addr: %s", BTCZRC20Addr.Hex())
//	BTCZRC20, err := zrc20.NewZRC20(BTCZRC20Addr, r.ZevmClient)
//	if err != nil {
//		panic(err)
//	}
//	balance, err := BTCZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
//	if err != nil {
//		panic(err)
//	}
//	if balance.Cmp(totalAmount) < 0 {
//		panic(fmt.Errorf("not enough balance in ZRC20 contract"))
//	}
//	// approve the ZRC20 contract to spend 1 BTC from the deployer address
//	{
//		// approve more to cover withdraw fee
//		tx, err := BTCZRC20.Approve(r.ZevmAuth, BTCZRC20Addr, totalAmount.Mul(totalAmount, big.NewInt(100)))
//		if err != nil {
//			panic(err)
//		}
//		receipt := config.MustWaitForTxReceipt(r.ZevmClient, tx, r.Logger)
//		r.Logger.Info("approve receipt: status %d", receipt.Status)
//		if receipt.Status != 1 {
//			panic(fmt.Errorf("approve receipt status is not 1"))
//		}
//	}
//	go func() {
//		for {
//			time.Sleep(3 * time.Second)
//			_, err = r.BtcRPCClient.GenerateToAddress(1, r.BTCDeployerAddress, nil)
//			if err != nil {
//				panic(err)
//			}
//		}
//	}()
//	// withdraw 0.1 BTC from ZRC20 to BTC address
//	for i := int64(0); i < repeat; i++ {
//		_, gasFee, err := BTCZRC20.WithdrawGasFee(&bind.CallOpts{})
//		if err != nil {
//			panic(err)
//		}
//		r.Logger.Info("withdraw gas fee: %d", gasFee)
//		tx, err := BTCZRC20.Withdraw(r.ZevmAuth, []byte(r.BTCDeployerAddress.EncodeAddress()), amount)
//		if err != nil {
//			panic(err)
//		}
//		receipt := config.MustWaitForTxReceipt(r.ZevmClient, tx, r.Logger)
//		r.Logger.Info("withdraw receipt: status %d", receipt.Status)
//		if receipt.Status != 1 {
//			panic(fmt.Errorf("withdraw receipt status is not 1"))
//		}
//		_, err = r.BtcRPCClient.GenerateToAddress(10, r.BTCDeployerAddress, nil)
//		if err != nil {
//			panic(err)
//		}
//		cctx := config.WaitCctxMinedByInTxHash(receipt.TxHash.Hex(), r.CctxClient, r.Logger)
//		outTxHash := cctx.GetCurrentOutTxParam().OutboundTxHash
//		hash, err := chainhash.NewHashFromStr(outTxHash)
//		if err != nil {
//			panic(err)
//		}
//
//		rawTx, err := r.BtcRPCClient.GetRawTransactionVerbose(hash)
//		if err != nil {
//			panic(err)
//		}
//		r.Logger.Info("raw tx:")
//		r.Logger.Info("  TxIn: %d", len(rawTx.Vin))
//		for idx, txIn := range rawTx.Vin {
//			r.Logger.Info("  TxIn %d:", idx)
//			r.Logger.Info("    TxID:Vout:  %s:%d", txIn.Txid, txIn.Vout)
//			r.Logger.Info("    ScriptSig: %s", txIn.ScriptSig.Hex)
//		}
//		r.Logger.Info("  TxOut: %d", len(rawTx.Vout))
//		for _, txOut := range rawTx.Vout {
//			r.Logger.Info("  TxOut %d:", txOut.N)
//			r.Logger.Info("    Value: %.8f", txOut.Value)
//			r.Logger.Info("    ScriptPubKey: %s", txOut.ScriptPubKey.Hex)
//		}
//	}
//}

// DepositBTCRefund ...
// TODO: define e2e test
// https://github.com/zeta-chain/node-private/issues/79
//func DepositBTCRefund(r *runner.E2ERunner) {
//	r.Logger.InfoLoud("Deposit BTC with invalid memo; should be refunded")
//	btc := r.BtcRPCClient
//	utxos, err := r.BtcRPCClient.ListUnspent()
//	if err != nil {
//		panic(err)
//	}
//	spendableAmount := 0.0
//	spendableUTXOs := 0
//	for _, utxo := range utxos {
//		if utxo.Spendable {
//			spendableAmount += utxo.Amount
//			spendableUTXOs++
//		}
//	}
//	r.Logger.Info("ListUnspent:")
//	r.Logger.Info("  spendableAmount: %f", spendableAmount)
//	r.Logger.Info("  spendableUTXOs: %d", spendableUTXOs)
//	r.Logger.Info("Now sending two txs to TSS address...")
//	_, err = r.SendToTSSFromDeployerToDeposit(r.BTCTSSAddress, 1.1, utxos[:2], btc, r.BTCDeployerAddress)
//	if err != nil {
//		panic(err)
//	}
//	_, err = r.SendToTSSFromDeployerToDeposit(r.BTCTSSAddress, 0.05, utxos[2:4], btc, r.BTCDeployerAddress)
//	if err != nil {
//		panic(err)
//	}
//
//	r.Logger.Info("testing if the deposit into BTC ZRC20 is successful...")
//
//	// check if the deposit is successful
//	initialBalance, err := r.BTCZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
//	if err != nil {
//		panic(err)
//	}
//	for {
//		time.Sleep(3 * time.Second)
//		balance, err := r.BTCZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
//		if err != nil {
//			panic(err)
//		}
//		diff := big.NewInt(0)
//		diff.Sub(balance, initialBalance)
//		if diff.Cmp(big.NewInt(1.15*btcutil.SatoshiPerBitcoin)) != 0 {
//			r.Logger.Info("waiting for BTC balance to show up in ZRC contract... current bal %d", balance)
//		} else {
//			r.Logger.Info("BTC balance is in ZRC contract! Success")
//			break
//		}
//	}
//}
