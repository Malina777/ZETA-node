package smoketests

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/erc20custody.sol"
	"github.com/zeta-chain/zetacore/contrib/localnet/orchestrator/smoketest/runner"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
)

func TestWhitelistERC20(sm *runner.SmokeTestRunner) {
	res, err := sm.ObserverClient.GetCoreParamsForChain(context.Background(), &observertypes.QueryGetCoreParamsForChainRequest{
		ChainId: int64(1337),
	})
	if err != nil {
		panic(err)
	}
	custodyAddr := ethcommon.HexToAddress(res.CoreParams.Erc20CustodyContractAddress)
	if custodyAddr == (ethcommon.Address{}) {
		panic("custody address is empty")
	}
	custody, err := erc20custody.NewERC20Custody(custodyAddr, sm.GoerliClient)
	if err != nil {
		panic(err)
	}
	iter, err := custody.FilterWhitelisted(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: context.Background(),
	}, []ethcommon.Address{})
	if err != nil {
		panic(err)
	}
	for iter.Next() {
		sm.Logger.Info("whitelisted: %s", iter.Event.Asset.Hex())
	}
}
