package e2etests

import (
	"fmt"
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	"golang.org/x/sync/errgroup"
)

// TestStressEtherDeposit tests the stressing deposit of ether
func TestStressEtherDeposit(r *runner.E2ERunner) {
	// number of deposits to perform
	numDeposits := 100

	r.Logger.Print("starting stress test of %d deposits", numDeposits)

	// create a wait group to wait for all the deposits to complete
	var eg errgroup.Group

	// send the deposits
	for i := 0; i < numDeposits; i++ {
		i := i
		hash := r.DepositEtherWithAmount(false, big.NewInt(100000))
		r.Logger.Print("index %d: starting deposit, tx hash: %s", i, hash.Hex())

		eg.Go(func() error {
			return MonitorEtherDeposit(r, hash, i, time.Now())
		})
	}

	// wait for all the deposits to complete
	if err := eg.Wait(); err != nil {
		panic(err)
	}

	r.Logger.Print("all deposits completed")
}

// MonitorEtherDeposit monitors the deposit of ether, returns once the deposit is complete
func MonitorEtherDeposit(r *runner.E2ERunner, hash ethcommon.Hash, index int, startTime time.Time) error {
	cctx := utils.WaitCctxMinedByInTxHash(r.Ctx, hash.Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
	if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
		return fmt.Errorf(
			"index %d: deposit cctx failed with status %s, message %s, cctx index %s",
			index,
			cctx.CctxStatus.Status,
			cctx.CctxStatus.StatusMessage,
			cctx.Index,
		)
	}
	timeToComplete := time.Now().Sub(startTime)
	r.Logger.Print("index %d: deposit cctx success in %s", index, timeToComplete.String())

	return nil
}
