package e2etests

import (
	"math/big"

	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
)

// TestDonationEther tests donation of ether to the tss address
func TestDonationEther(r *runner.E2ERunner) {
	txDonation, err := r.SendEther(r.TSSAddress, big.NewInt(100000000000000000), []byte(common.DonationMessage))
	if err != nil {
		panic(err)
	}
	r.Logger.EVMTransaction(*txDonation, "donation")

	// check contract deployment receipt
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.GoerliClient, txDonation, r.Logger, r.ReceiptTimeout)
	r.Logger.EVMReceipt(*receipt, "donation")
	if receipt.Status != 1 {
		panic("donation tx failed")
	}
}
