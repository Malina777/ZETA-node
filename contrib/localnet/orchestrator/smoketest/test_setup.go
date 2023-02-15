package main

import (
	"context"
	"fmt"
	"github.com/zeta-chain/zetacore/contrib/localnet/orchestrator/smoketest/contracts/testdapp"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeta-chain/zetacore/contracts/evm/erc20custody"
	"github.com/zeta-chain/zetacore/contracts/evm/zetaconnectoreth"
	"github.com/zeta-chain/zetacore/contracts/evm/zetaeth"
	contracts "github.com/zeta-chain/zetacore/contracts/zevm"
	"github.com/zeta-chain/zetacore/contrib/localnet/orchestrator/smoketest/contracts/erc20"
	fungibletypes "github.com/zeta-chain/zetacore/x/fungible/types"
)

func (sm *SmokeTest) TestSetupZetaTokenAndConnectorContracts() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("test finishes in %s\n", time.Since(startTime))
	}()

	goerliClient := sm.goerliClient
	chainid, err := goerliClient.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	deployerPrivkey, err := crypto.HexToECDSA(DeployerPrivateKey)
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(deployerPrivkey, chainid)
	if err != nil {
		panic(err)
	}

	LoudPrintf("Deploy ZetaETH ConnectorETH ERC20Custody USDT\n")
	if err := CheckNonce(goerliClient, DeployerAddress, 0); err != nil {
		panic(err)
	}
	zetaEthAddr, tx, ZetaEth, err := zetaeth.DeployZetaEth(auth, goerliClient, big.NewInt(21_000_000_000))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ZetaEth contract address: %s, tx hash: %s\n", zetaEthAddr.Hex(), tx.Hash().Hex())
	receipt := MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("ZetaEth contract receipt: contract address %s, status %d\n", receipt.ContractAddress, receipt.Status)
	sm.ZetaEth = ZetaEth
	sm.ZetaEthAddr = zetaEthAddr

	if err := CheckNonce(goerliClient, DeployerAddress, 1); err != nil {
		panic(err)
	}
	connectorEthAddr, tx, ConnectorEth, err := zetaconnectoreth.DeployZetaConnectorEth(auth, goerliClient, zetaEthAddr,
		TSSAddress, DeployerAddress, DeployerAddress)
	if err != nil {
		panic(err)
	}
	receipt = MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("ZetaConnectorEth contract address: %s, tx hash: %s\n", connectorEthAddr.Hex(), tx.Hash().Hex())
	fmt.Printf("ZetaConnectorEth contract receipt: contract address %s, status %d\n", receipt.ContractAddress, receipt.Status)
	sm.ConnectorEth = ConnectorEth
	sm.ConnectorEthAddr = connectorEthAddr

	fungibleClient := sm.fungibleClient
	fmt.Printf("Deploying ERC20Custody contract\n")
	if err := CheckNonce(goerliClient, DeployerAddress, 2); err != nil {
		panic(err)
	}
	erc20CustodyAddr, tx, ERC20Custody, err := erc20custody.DeployERC20Custody(auth, goerliClient, DeployerAddress, DeployerAddress, big.NewInt(0), ethcommon.HexToAddress("0x"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ERC20Custody contract address: %s, tx hash: %s\n", erc20CustodyAddr.Hex(), tx.Hash().Hex())
	receipt = MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("ERC20Custody contract receipt: contract address %s, status %d\n", receipt.ContractAddress, receipt.Status)
	if erc20CustodyAddr != ethcommon.HexToAddress(ERC20CustodyAddr) {
		panic("ERC20Custody contract address mismatch! check order of tx")
	}
	sm.ERC20CustodyAddr = erc20CustodyAddr
	sm.ERC20Custody = ERC20Custody

	fmt.Printf("Deploying USDT contract\n")
	if err := CheckNonce(goerliClient, DeployerAddress, 3); err != nil {
		panic(err)
	}
	usdtAddr, tx, _, err := erc20.DeployUSDT(auth, goerliClient, "USDT", "USDT", 6)
	if err != nil {
		panic(err)
	}
	fmt.Printf("USDT contract address: %s, tx hash: %s\n", usdtAddr.Hex(), tx.Hash().Hex())
	receipt = MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("USDT contract receipt: contract address %s, status %d\n", receipt.ContractAddress, receipt.Status)
	if receipt.ContractAddress != ethcommon.HexToAddress(USDTERC20Addr) {
		panic("USDT contract address mismatch! check order of tx")
	}
	fmt.Printf("Step 6: Whitelist USDT\n")
	tx, err = ERC20Custody.Whitelist(auth, usdtAddr)
	if err != nil {
		panic(err)
	}
	receipt = MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("Whitelist receipt tx hash: %s\n", tx.Hash().Hex())

	fmt.Printf("Step 7: Set TSS address\n")
	tx, err = ERC20Custody.UpdateTSSAddress(auth, TSSAddress)
	if err != nil {
		panic(err)
	}
	receipt = MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("TSS set receipt tx hash: %s\n", tx.Hash().Hex())

	fmt.Printf("Checking foreign coins...\n")
	res, err := fungibleClient.ForeignCoinsAll(context.Background(), &fungibletypes.QueryAllForeignCoinsRequest{})
	if err != nil {
		panic(err)
	}
	found := false
	zrc20addr := ""
	for _, fcoin := range res.ForeignCoins {
		if ethcommon.HexToAddress(fcoin.Asset) == usdtAddr {
			found = true
			zrc20addr = fcoin.Zrc20ContractAddress
		}
	}
	if !found {
		fmt.Printf("foreign coins: %v", res.ForeignCoins)
		panic(fmt.Sprintf("fungible module does not have foreign coin that represent USDT ERC20 %s", usdtAddr))
	}
	fmt.Printf("USDT ZRC20 Address: %s\n", zrc20addr)
	if HexToAddress(zrc20addr) != HexToAddress(USDTZRC20Addr) {
		panic("mismatch of foreign coin USDT ZRC20 and the USDTZRC20Addr constant in smoketest")
	}
	sm.USDTZRC20Addr = ethcommon.HexToAddress(zrc20addr)
	sm.USDTZRC20, err = contracts.NewZRC20(sm.USDTZRC20Addr, sm.zevmClient)
	if err != nil {
		panic(err)
	}

	USDT, err := erc20.NewUSDT(usdtAddr, goerliClient)
	if err != nil {
		panic(err)
	}
	sm.USDTERC20 = USDT
	sm.USDTERC20Addr = usdtAddr
	sm.UniswapV2FactoryAddr = ethcommon.HexToAddress(UniswapV2FactoryAddr)
	sm.UniswapV2Factory, err = contracts.NewUniswapV2Factory(sm.UniswapV2FactoryAddr, sm.zevmClient)
	if err != nil {
		panic(err)
	}
	sm.UniswapV2RouterAddr = ethcommon.HexToAddress(UniswapV2RouterAddr)
	sm.UniswapV2Router, err = contracts.NewUniswapV2Router02(sm.UniswapV2RouterAddr, sm.zevmClient)
	if err != nil {
		panic(err)
	}

	// deploy TestDApp contract
	//auth.GasLimit = 1_000_000
	appAddr, tx, _, err := testdapp.DeployTestDApp(auth, goerliClient, sm.ConnectorEthAddr, sm.ZetaEthAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDApp contract address: %s, tx hash: %s\n", appAddr.Hex(), tx.Hash().Hex())
	receipt = MustWaitForTxReceipt(goerliClient, tx)
	fmt.Printf("TestDApp contract receipt: contract address %s, status %d; used gas %d\n", receipt.ContractAddress, receipt.Status, receipt.GasUsed)
	dapp, err := testdapp.NewTestDApp(receipt.ContractAddress, goerliClient)
	if err != nil {
		panic(err)
	}
	{
		time.Sleep(10 * time.Second)
		code, err := sm.goerliClient.CodeAt(context.Background(), receipt.ContractAddress, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("TestDApp contract code: len %d\n", len(code))
		if len(code) == 0 {
			panic("TestDApp contract code is empty")
		}
		res, err := dapp.Connector(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		if res != sm.ConnectorEthAddr {
			panic("mismatch of TestDApp connector address")
		}
	}
	sm.TestDAppAddr = receipt.ContractAddress

}
