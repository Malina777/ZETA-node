package keeper_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/cmd/zetacored/config"
	"github.com/zeta-chain/zetacore/common"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
	crosschainkeeper "github.com/zeta-chain/zetacore/x/crosschain/keeper"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	fungiblekeeper "github.com/zeta-chain/zetacore/x/fungible/keeper"
	fungibletypes "github.com/zeta-chain/zetacore/x/fungible/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
)

func TestParseZRC20WithdrawalEvent(t *testing.T) {
	t.Run("unable to parse an event with an invalid address in event log", func(t *testing.T) {
		for i, log := range GetInvalidZRC20WithdrawToExternal(t).Logs {
			event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*log)
			if i < 3 {
				require.ErrorContains(t, err, "event signature mismatch")
				require.Nil(t, event)
				continue
			}
			require.NoError(t, err)
			require.NotNil(t, event)
			require.Equal(t, "1EYVvXLusCxtVuEwoYvWRyN5EZTXwPVvo3", string(event.To))
		}
	})
	t.Run("successfully parse event for a valid BTC withdrawal", func(t *testing.T) {
		for i, log := range GetValidZRC20WithdrawToBTC(t).Logs {
			event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*log)
			if i < 3 {
				require.ErrorContains(t, err, "event signature mismatch")
				require.Nil(t, event)
				continue
			}
			require.NoError(t, err)
			require.NotNil(t, event)
			require.Equal(t, "0x33EaD83db0D0c682B05ead61E8d8f481Bb1B4933", event.From.Hex())
			require.Equal(t, "bc1qysd4sp9q8my59ul9wsf5rvs9p387hf8vfwatzu", string(event.To))
		}
	})
	t.Run("successfully parse valid event for ETH withdrawal", func(t *testing.T) {
		for i, log := range GetValidZrc20WithdrawToETH(t).Logs {
			event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*log)
			if i != 11 {
				require.ErrorContains(t, err, "event signature mismatch")
				require.Nil(t, event)
				continue
			}
			require.NoError(t, err)
			require.NotNil(t, event)
			require.Equal(t, "0x5daBFdd153Aaab4a970fD953DcFEEE8BF6Bb946E", ethcommon.BytesToAddress(event.To).String())
			require.Equal(t, "0x8E0f8E7E9E121403e72151d00F4937eACB2D9Ef3", event.From.Hex())
		}
	})
	t.Run("failed to parse event with a valid address but no topic present", func(t *testing.T) {
		for _, log := range GetValidZRC20WithdrawToBTC(t).Logs {
			log.Topics = nil
			event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*log)
			require.ErrorContains(t, err, "invalid log - no topics")
			require.Nil(t, event)
		}
	})
}
func TestValidateZrc20WithdrawEvent(t *testing.T) {
	t.Run("successfully validate a valid event", func(t *testing.T) {
		btcMainNetWithdrawalEvent, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZRC20WithdrawToBTC(t).Logs[3])
		require.NoError(t, err)
		err = crosschainkeeper.ValidateZrc20WithdrawEvent(btcMainNetWithdrawalEvent, common.BtcMainnetChain().ChainId)
		require.NoError(t, err)
	})
	t.Run("unable to validate a event with an invalid amount", func(t *testing.T) {
		btcMainNetWithdrawalEvent, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZRC20WithdrawToBTC(t).Logs[3])
		require.NoError(t, err)
		btcMainNetWithdrawalEvent.Value = big.NewInt(0)
		err = crosschainkeeper.ValidateZrc20WithdrawEvent(btcMainNetWithdrawalEvent, common.BtcMainnetChain().ChainId)
		require.ErrorContains(t, err, "ParseZRC20WithdrawalEvent: invalid amount")
	})
	t.Run("unable to validate a event with an invalid chain ID", func(t *testing.T) {
		btcMainNetWithdrawalEvent, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZRC20WithdrawToBTC(t).Logs[3])
		require.NoError(t, err)
		err = crosschainkeeper.ValidateZrc20WithdrawEvent(btcMainNetWithdrawalEvent, common.BtcTestNetChain().ChainId)
		require.ErrorContains(t, err, "address is not for network testnet3")
	})
	t.Run("unable to validate a event with an invalid address type", func(t *testing.T) {
		btcMainNetWithdrawalEvent, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZRC20WithdrawToBTC(t).Logs[3])
		require.NoError(t, err)
		btcMainNetWithdrawalEvent.To = []byte("1EYVvXLusCxtVuEwoYvWRyN5EZTXwPVvo3")
		err = crosschainkeeper.ValidateZrc20WithdrawEvent(btcMainNetWithdrawalEvent, common.BtcTestNetChain().ChainId)
		require.ErrorContains(t, err, "decode address failed: unknown address type")
	})
}

func TestKeeper_ProcessZRC20WithdrawalEvent(t *testing.T) {
	t.Run("successfully process ZRC20Withdrawal to BTC chain", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZRC20WithdrawToBTC(t).Logs[3])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 1)
		require.Equal(t, "bc1qysd4sp9q8my59ul9wsf5rvs9p387hf8vfwatzu", cctxList[0].GetCurrentOutTxParam().Receiver)
		require.Equal(t, emittingContract.Hex(), cctxList[0].InboundTxParams.Sender)
		require.Equal(t, txOrigin.Hex(), cctxList[0].InboundTxParams.TxOrigin)
	})
	t.Run("successfully process ZRC20Withdrawal to ETH chain", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 1)
		require.Equal(t, "0x5daBFdd153Aaab4a970fD953DcFEEE8BF6Bb946E", cctxList[0].GetCurrentOutTxParam().Receiver)
		require.Equal(t, emittingContract.Hex(), cctxList[0].InboundTxParams.Sender)
		require.Equal(t, txOrigin.Hex(), cctxList[0].InboundTxParams.TxOrigin)
	})
	t.Run("unable to process ZRC20Withdrawal if foreign coin is not found", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "cannot find foreign coin with emittingContract address")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
	t.Run("unable to process ZRC20Withdrawal if receiver chain is not supported", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "chain not supported")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
	t.Run("unable to process ZRC20Withdrawal if zeta chainID is not correctly set", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()
		ctx = ctx.WithChainID("test_21-1")

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "failed to convert chainID: chain 21 not found")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
	t.Run("unable to process ZRC20Withdrawal if to address is not in correct format", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		event.To = ethcommon.Address{}.Bytes()
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "cannot encode address")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
	t.Run("unable to process ZRC20Withdrawal if gaslimit not set on zrc20 contract", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeperWithMocks(t, keepertest.CrosschainMockOptions{
			UseFungibleMock: true,
		})
		fungibleMock := keepertest.GetCrosschainFungibleMock(t, k)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()
		fc, _ := zk.FungibleKeeper.GetForeignCoins(ctx, zrc20.Hex())

		fungibleMock.On("GetForeignCoins", mock.Anything, mock.Anything).Return(fc, true)
		fungibleMock.On("QueryGasLimit", mock.Anything, mock.Anything).Return(big.NewInt(0), fmt.Errorf("error querying gas limit"))
		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "error querying gas limit")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
	t.Run("unable to process ZRC20Withdrawal if gasprice is not set in crosschain keeper", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		k.RemoveGasPrice(ctx, strconv.FormatInt(chainID, 10))

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "gasprice not found")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
	t.Run("unable to process ZRC20Withdrawal if process cctx fails", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		zk.ObserverKeeper.SetChainNonces(ctx, observertypes.ChainNonces{
			Index:   chain.ChainName.String(),
			ChainId: chain.ChainId,
			Nonce:   1,
		})

		event, err := crosschainkeeper.ParseZRC20WithdrawalEvent(*GetValidZrc20WithdrawToETH(t).Logs[11])
		require.NoError(t, err)
		zrc20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "ethereum", "ETH")
		event.Raw.Address = zrc20
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZRC20WithdrawalEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "ProcessWithdrawalEvent: update nonce failed")
		require.Empty(t, k.GetAllCrossChainTx(ctx))
	})
}

func TestKeeper_ParseZetaSentEvent(t *testing.T) {
	t.Run("successfully parse a valid event", func(t *testing.T) {
		logs := GetValidZetaSentDestinationExternal(t).Logs
		for i, log := range logs {
			connector := log.Address
			event, err := crosschainkeeper.ParseZetaSentEvent(*log, connector)
			if i < 4 {
				require.ErrorContains(t, err, "event signature mismatch")
				require.Nil(t, event)
				continue
			}
			require.Equal(t, common.EthChain().ChainId, event.DestinationChainId.Int64())
			require.Equal(t, "70000000000000000000", event.ZetaValueAndGas.String())
			require.Equal(t, "0x60983881bdf302dcfa96603A58274D15D5966209", event.SourceTxOriginAddress.String())
			require.Equal(t, "0xF0a3F93Ed1B126142E61423F9546bf1323Ff82DF", event.ZetaTxSenderAddress.String())
		}
	})
	t.Run("unable to parse if topics field is empty", func(t *testing.T) {
		logs := GetValidZetaSentDestinationExternal(t).Logs
		for _, log := range logs {
			connector := log.Address
			log.Topics = nil
			event, err := crosschainkeeper.ParseZetaSentEvent(*log, connector)
			require.ErrorContains(t, err, "ParseZetaSentEvent: invalid log - no topics")
			require.Nil(t, event)
		}
	})
	t.Run("unable to parse if connector address does not match", func(t *testing.T) {
		logs := GetValidZetaSentDestinationExternal(t).Logs
		for i, log := range logs {
			event, err := crosschainkeeper.ParseZetaSentEvent(*log, sample.EthAddress())
			if i < 4 {
				require.ErrorContains(t, err, "event signature mismatch")
				require.Nil(t, event)
				continue
			}
			require.ErrorContains(t, err, "does not match connectorZEVM")
			require.Nil(t, event)
		}
	})
}
func TestKeeper_ProcessZetaSentEvent(t *testing.T) {
	t.Run("successfully process ZetaSentEvent", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)

		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		SetupStateForProcessLogsZetaSent(t, ctx, k, zk, sdkk, chain)
		amount, ok := sdkmath.NewIntFromString("20000000000000000000000")
		require.True(t, ok)
		err := sdkk.BankKeeper.MintCoins(ctx, fungibletypes.ModuleName, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, amount)))
		require.NoError(t, err)

		event, err := crosschainkeeper.ParseZetaSentEvent(*GetValidZetaSentDestinationExternal(t).Logs[4], GetValidZetaSentDestinationExternal(t).Logs[4].Address)
		require.NoError(t, err)
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZetaSentEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 1)
		require.Equal(t, strings.Compare("0x60983881bdf302dcfa96603a58274d15d5966209", cctxList[0].GetCurrentOutTxParam().Receiver), 0)
		require.Equal(t, common.EthChain().ChainId, cctxList[0].GetCurrentOutTxParam().ReceiverChainId)
		require.Equal(t, emittingContract.Hex(), cctxList[0].InboundTxParams.Sender)
		require.Equal(t, txOrigin.Hex(), cctxList[0].InboundTxParams.TxOrigin)
	})
	t.Run("unable to process ZetaSentEvent if fungible module does not have enough balance", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		SetupStateForProcessLogsZetaSent(t, ctx, k, zk, sdkk, chain)

		event, err := crosschainkeeper.ParseZetaSentEvent(*GetValidZetaSentDestinationExternal(t).Logs[4], GetValidZetaSentDestinationExternal(t).Logs[4].Address)
		require.NoError(t, err)
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZetaSentEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "ProcessZetaSentEvent: failed to burn coins from fungible")
	})
	t.Run("unable to process ZetaSentEvent if receiver chain is not supported", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		SetupStateForProcessLogsZetaSent(t, ctx, k, zk, sdkk, chain)

		amount, ok := sdkmath.NewIntFromString("20000000000000000000000")
		require.True(t, ok)
		err := sdkk.BankKeeper.MintCoins(ctx, fungibletypes.ModuleName, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, amount)))
		require.NoError(t, err)

		event, err := crosschainkeeper.ParseZetaSentEvent(*GetValidZetaSentDestinationExternal(t).Logs[4], GetValidZetaSentDestinationExternal(t).Logs[4].Address)
		require.NoError(t, err)
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()
		err = k.ProcessZetaSentEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "chain not supported")
	})
	t.Run("unable to process ZetaSentEvent if zetachain chain id not correctly set in context", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		SetupStateForProcessLogsZetaSent(t, ctx, k, zk, sdkk, chain)

		amount, ok := sdkmath.NewIntFromString("20000000000000000000000")
		require.True(t, ok)
		err := sdkk.BankKeeper.MintCoins(ctx, fungibletypes.ModuleName, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, amount)))
		require.NoError(t, err)

		event, err := crosschainkeeper.ParseZetaSentEvent(*GetValidZetaSentDestinationExternal(t).Logs[4], GetValidZetaSentDestinationExternal(t).Logs[4].Address)
		require.NoError(t, err)
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()
		ctx = ctx.WithChainID("test-21-1")
		err = k.ProcessZetaSentEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "ProcessZetaSentEvent: failed to convert chainID")
	})

	t.Run("unable to process ZetaSentEvent if gas pay fails", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		amount, ok := sdkmath.NewIntFromString("20000000000000000000000")
		require.True(t, ok)
		err := sdkk.BankKeeper.MintCoins(ctx, fungibletypes.ModuleName, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, amount)))
		require.NoError(t, err)
		event, err := crosschainkeeper.ParseZetaSentEvent(*GetValidZetaSentDestinationExternal(t).Logs[4], GetValidZetaSentDestinationExternal(t).Logs[4].Address)
		require.NoError(t, err)
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()

		err = k.ProcessZetaSentEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "ProcessWithdrawalEvent: pay gas failed")
	})
	t.Run("unable to process ZetaSentEvent if process cctx fails", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)

		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		SetupStateForProcessLogsZetaSent(t, ctx, k, zk, sdkk, chain)

		zk.ObserverKeeper.SetChainNonces(ctx, observertypes.ChainNonces{
			Index:   chain.ChainName.String(),
			ChainId: chain.ChainId,
			Nonce:   1,
		})
		amount, ok := sdkmath.NewIntFromString("20000000000000000000000")
		require.True(t, ok)
		err := sdkk.BankKeeper.MintCoins(ctx, fungibletypes.ModuleName, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, amount)))
		require.NoError(t, err)

		event, err := crosschainkeeper.ParseZetaSentEvent(*GetValidZetaSentDestinationExternal(t).Logs[4], GetValidZetaSentDestinationExternal(t).Logs[4].Address)
		require.NoError(t, err)
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()
		tss := sample.Tss()
		err = k.ProcessZetaSentEvent(ctx, event, emittingContract, txOrigin.Hex(), tss)
		require.ErrorContains(t, err, "ProcessWithdrawalEvent: update nonce failed")
	})
}

func TestKeeper_ProcessLogs(t *testing.T) {
	t.Run("successfully parse and process ZRC20Withdrawal to BTC chain", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		block := GetValidZRC20WithdrawToBTC(t)
		gasZRC20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		for _, log := range block.Logs {
			log.Address = gasZRC20
		}
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()

		err := k.ProcessLogs(ctx, block.Logs, emittingContract, txOrigin.Hex())
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 1)
		require.Equal(t, "bc1qysd4sp9q8my59ul9wsf5rvs9p387hf8vfwatzu", cctxList[0].GetCurrentOutTxParam().Receiver)
		require.Equal(t, emittingContract.Hex(), cctxList[0].InboundTxParams.Sender)
		require.Equal(t, txOrigin.Hex(), cctxList[0].InboundTxParams.TxOrigin)
	})
	t.Run("successfully parse and process ZetaSentEvent", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.EthChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)
		SetupStateForProcessLogsZetaSent(t, ctx, k, zk, sdkk, chain)

		amount, ok := sdkmath.NewIntFromString("20000000000000000000000")
		require.True(t, ok)
		err := sdkk.BankKeeper.MintCoins(ctx, fungibletypes.ModuleName, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, amount)))
		require.NoError(t, err)
		block := GetValidZetaSentDestinationExternal(t)
		system, found := zk.FungibleKeeper.GetSystemContract(ctx)
		require.True(t, found)
		for _, log := range block.Logs {
			log.Address = ethcommon.HexToAddress(system.ConnectorZevm)
		}
		emittingContract := sample.EthAddress()
		txOrigin := sample.EthAddress()

		err = k.ProcessLogs(ctx, block.Logs, emittingContract, txOrigin.Hex())
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 1)
		require.Equal(t, strings.Compare("0x60983881bdf302dcfa96603a58274d15d5966209", cctxList[0].GetCurrentOutTxParam().Receiver), 0)
		require.Equal(t, common.EthChain().ChainId, cctxList[0].GetCurrentOutTxParam().ReceiverChainId)
		require.Equal(t, emittingContract.Hex(), cctxList[0].InboundTxParams.Sender)
		require.Equal(t, txOrigin.Hex(), cctxList[0].InboundTxParams.TxOrigin)
	})
	t.Run("unable to process logs if system contract not found", func(t *testing.T) {
		k, ctx, _, _ := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		err := k.ProcessLogs(ctx, GetValidZRC20WithdrawToBTC(t).Logs, sample.EthAddress(), "")
		require.ErrorContains(t, err, "cannot find system contract")
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 0)
	})
	t.Run("no cctx created for logs containing no events", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		block := GetValidZRC20WithdrawToBTC(t)
		gasZRC20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		for _, log := range block.Logs {
			log.Address = gasZRC20
		}
		block.Logs = block.Logs[:3]

		err := k.ProcessLogs(ctx, block.Logs, sample.EthAddress(), "")
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 0)
	})
	t.Run("no cctx created  for logs containing proper event but not emitted from a known ZRC20 contract", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)
		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		block := GetValidZRC20WithdrawToBTC(t)
		setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		for _, log := range block.Logs {
			log.Address = sample.EthAddress()
		}

		err := k.ProcessLogs(ctx, block.Logs, sample.EthAddress(), "")
		require.NoError(t, err)
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 0)
	})
	t.Run("no cctx created for valid logs if Inbound is disabled", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		block := GetValidZRC20WithdrawToBTC(t)
		gasZRC20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		for _, log := range block.Logs {
			log.Address = gasZRC20
		}
		zk.ObserverKeeper.SetCrosschainFlags(ctx, observertypes.CrosschainFlags{
			IsInboundEnabled: false,
		})

		err := k.ProcessLogs(ctx, block.Logs, sample.EthAddress(), "")
		require.ErrorContains(t, err, observertypes.ErrInboundDisabled.Error())
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 0)
	})
	t.Run("error returned for invalid event data", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		block := GetInvalidZRC20WithdrawToExternal(t)
		gasZRC20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		for _, log := range block.Logs {
			log.Address = gasZRC20
		}

		err := k.ProcessLogs(ctx, block.Logs, sample.EthAddress(), "")
		require.ErrorContains(t, err, "ParseZRC20WithdrawalEvent: invalid address")
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 0)
	})
	t.Run("error returned if unable to process an event", func(t *testing.T) {
		k, ctx, sdkk, zk := keepertest.CrosschainKeeper(t)
		k.GetAuthKeeper().GetModuleAccount(ctx, fungibletypes.ModuleName)

		chain := common.BtcMainnetChain()
		chainID := chain.ChainId
		setSupportedChain(ctx, zk, chainID)
		SetupStateForProcessLogs(t, ctx, k, zk, sdkk, chain)

		block := GetValidZRC20WithdrawToBTC(t)
		gasZRC20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chainID, "bitcoin", "BTC")
		for _, log := range block.Logs {
			log.Address = gasZRC20
		}
		ctx = ctx.WithChainID("test-21-1")

		err := k.ProcessLogs(ctx, block.Logs, sample.EthAddress(), "")
		require.ErrorContains(t, err, "ProcessZRC20WithdrawalEvent: failed to convert chainID")
		cctxList := k.GetAllCrossChainTx(ctx)
		require.Len(t, cctxList, 0)
	})
}

func SetupStateForProcessLogsZetaSent(t *testing.T, ctx sdk.Context, k *crosschainkeeper.Keeper, zk keepertest.ZetaKeepers, sdkk keepertest.SDKKeepers, chain common.Chain) {
	admin := sample.AccAddress()
	setAdminPolicies(ctx, zk, admin)

	assetAddress := sample.EthAddress().String()
	gasZRC20 := setupGasCoin(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper, chain.ChainId, "ethereum", "ETH")
	zrc20Addr := deployZRC20(
		t,
		ctx,
		zk.FungibleKeeper,
		sdkk.EvmKeeper,
		chain.ChainId,
		"ethereum",
		assetAddress,
		"ETH",
	)
	fungibleMsgServer := fungiblekeeper.NewMsgServerImpl(*zk.FungibleKeeper)
	_, err := fungibleMsgServer.UpdateZRC20WithdrawFee(
		sdk.UnwrapSDKContext(ctx),
		fungibletypes.NewMsgUpdateZRC20WithdrawFee(admin, gasZRC20.String(), sdk.NewUint(withdrawFee), sdkmath.Uint{}),
	)
	require.NoError(t, err)
	k.SetGasPrice(ctx, crosschaintypes.GasPrice{
		ChainId:     chain.ChainId,
		MedianIndex: 0,
		Prices:      []uint64{gasPrice},
	})
	setupZRC20Pool(
		t,
		ctx,
		zk.FungibleKeeper,
		sdkk.BankKeeper,
		zrc20Addr,
	)
}
func SetupStateForProcessLogs(t *testing.T, ctx sdk.Context, k *crosschainkeeper.Keeper, zk keepertest.ZetaKeepers, sdkk keepertest.SDKKeepers, chain common.Chain) {

	deploySystemContracts(t, ctx, zk.FungibleKeeper, sdkk.EvmKeeper)
	tss := sample.Tss()
	zk.ObserverKeeper.SetTSS(ctx, tss)
	k.SetGasPrice(ctx, crosschaintypes.GasPrice{
		ChainId: chain.ChainId,
		Prices:  []uint64{100},
	})

	zk.ObserverKeeper.SetChainNonces(ctx, observertypes.ChainNonces{
		Index:   chain.ChainName.String(),
		ChainId: chain.ChainId,
		Nonce:   0,
	})
	zk.ObserverKeeper.SetPendingNonces(ctx, observertypes.PendingNonces{
		NonceLow:  0,
		NonceHigh: 0,
		ChainId:   chain.ChainId,
		Tss:       tss.TssPubkey,
	})
}

// receiver is 1EYVvXLusCxtVuEwoYvWRyN5EZTXwPVvo3
func GetInvalidZRC20WithdrawToExternal(t *testing.T) (receipt ethtypes.Receipt) {
	block := "{\n  \"type\": \"0x2\",\n  \"root\": \"0x\",\n  \"status\": \"0x1\",\n  \"cumulativeGasUsed\": \"0x4e7a38\",\n  \"logsBloom\": \"0x00000000000000000000010000020000000000000000000000000000000000020000000100000000000000000000000080000000000000000000000400200000200000000002000000000008000000000000000000000000000000000000000000000000020000000000000000800800000040000000000000000010000000000000000000000000000000000000000000000000000004000000000000000000020000000000000000000000000000000000000000000000000000000000010000000002000000000000000000000000000000000000000000000000000020000010000000000000000001000000000000000000040200000000000000000000\",\n  \"logs\": [\n    {\n      \"address\": \"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\n      \"topics\": [\n        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n        \"0x000000000000000000000000313e74f7755afbae4f90e02ca49f8f09ff934a37\",\n        \"0x000000000000000000000000735b14bb79463307aacbed86daf3322b1e6226ab\"\n      ],\n      \"data\": \"0x0000000000000000000000000000000000000000000000000000000000003790\",\n      \"blockNumber\": \"0x1a2ad3\",\n      \"transactionHash\": \"0x81126c18c7ca7d1fb7ded6644a87802e91bf52154ee4af7a5b379354e24fb6e0\",\n      \"transactionIndex\": \"0x10\",\n      \"blockHash\": \"0x5cb338544f64a226f4bfccb7a8d977f861c13ad73f7dd4317b66b00dd95de51c\",\n      \"logIndex\": \"0x46\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\n      \"topics\": [\n        \"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\n        \"0x000000000000000000000000313e74f7755afbae4f90e02ca49f8f09ff934a37\",\n        \"0x00000000000000000000000013a0c5930c028511dc02665e7285134b6d11a5f4\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000000000000006a1217\",\n      \"blockNumber\": \"0x1a2ad3\",\n      \"transactionHash\": \"0x81126c18c7ca7d1fb7ded6644a87802e91bf52154ee4af7a5b379354e24fb6e0\",\n      \"transactionIndex\": \"0x10\",\n      \"blockHash\": \"0x5cb338544f64a226f4bfccb7a8d977f861c13ad73f7dd4317b66b00dd95de51c\",\n      \"logIndex\": \"0x47\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\n      \"topics\": [\n        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n        \"0x000000000000000000000000313e74f7755afbae4f90e02ca49f8f09ff934a37\",\n        \"0x0000000000000000000000000000000000000000000000000000000000000000\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000000000000006a0c70\",\n      \"blockNumber\": \"0x1a2ad3\",\n      \"transactionHash\": \"0x81126c18c7ca7d1fb7ded6644a87802e91bf52154ee4af7a5b379354e24fb6e0\",\n      \"transactionIndex\": \"0x10\",\n      \"blockHash\": \"0x5cb338544f64a226f4bfccb7a8d977f861c13ad73f7dd4317b66b00dd95de51c\",\n      \"logIndex\": \"0x48\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\n      \"topics\": [\n        \"0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955\",\n        \"0x000000000000000000000000313e74f7755afbae4f90e02ca49f8f09ff934a37\"\n      ],\n      \"data\": \"0x000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000006a0c700000000000000000000000000000000000000000000000000000000000003790000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000223145595676584c7573437874567545776f59765752794e35455a5458775056766f33000000000000000000000000000000000000000000000000000000000000\",\n      \"blockNumber\": \"0x1a2ad3\",\n      \"transactionHash\": \"0x81126c18c7ca7d1fb7ded6644a87802e91bf52154ee4af7a5b379354e24fb6e0\",\n      \"transactionIndex\": \"0x10\",\n      \"blockHash\": \"0x5cb338544f64a226f4bfccb7a8d977f861c13ad73f7dd4317b66b00dd95de51c\",\n      \"logIndex\": \"0x49\",\n      \"removed\": false\n    }\n  ],\n  \"transactionHash\": \"0x81126c18c7ca7d1fb7ded6644a87802e91bf52154ee4af7a5b379354e24fb6e0\",\n  \"contractAddress\": \"0x0000000000000000000000000000000000000000\",\n  \"gasUsed\": \"0x12521\",\n  \"blockHash\": \"0x5cb338544f64a226f4bfccb7a8d977f861c13ad73f7dd4317b66b00dd95de51c\",\n  \"blockNumber\": \"0x1a2ad3\",\n  \"transactionIndex\": \"0x10\"\n}\n"
	err := json.Unmarshal([]byte(block), &receipt)
	require.NoError(t, err)
	return
}

func GetValidZrc20WithdrawToETH(t *testing.T) (receipt ethtypes.Receipt) {
	block := "{\n  \"type\": \"0x2\",\n  \"root\": \"0x\",\n  \"status\": \"0x1\",\n  \"cumulativeGasUsed\": \"0xdbedca\",\n  \"logsBloom\": \"0x00200000001000000000000088020001000001000000000000000000000000000000020100000000000000000000000080000000000000000000000400640000000000000000000008000008020000200000000000000002000000008000000000000000020000000200000000800801000000080000000000000010000000000000000000000000000000000000001000000001000004080001404000000000028002000000000000000040000000000000000000000000000000000000000000000002000000000000008000000000000000800800001000000002000021000010000100000000000010800400000000020000000100400880000000004000\",\n  \"logs\": [\n    {\n      \"address\": \"0x3f641963f3d9adf82d890fd8142313dcec807ba5\",\n      \"topics\": [\n        \"0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000045400a8fd5330000\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x57\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\n      \"topics\": [\n        \"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000001ac7c4159f72b90000\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x58\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\n      \"topics\": [\n        \"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n        \"0x0000000000000000000000002ca7d64a7efe2d62a725e2b35cf7230d6677ffee\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000001ac7c4159f72b90000\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x59\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\n      \"topics\": [\n        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n        \"0x00000000000000000000000016ef1b018026e389fda93c1e993e987cf6e852e7\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000001ac7c4159f72b90000\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x5a\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0xd97b1de3619ed2c6beb3860147e30ca8a7dc9891\",\n      \"topics\": [\n        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n        \"0x00000000000000000000000016ef1b018026e389fda93c1e993e987cf6e852e7\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000002e640d76638740f\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x5b\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x16ef1b018026e389fda93c1e993e987cf6e852e7\",\n      \"topics\": [\n        \"0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1\"\n      ],\n      \"data\": \"0x000000000000000000000000000000000000000000000b3f1da425061770a11600000000000000000000000000000000000000000000000135be3952e251aa40\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x5c\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x16ef1b018026e389fda93c1e993e987cf6e852e7\",\n      \"topics\": [\n        \"0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822\",\n        \"0x0000000000000000000000002ca7d64a7efe2d62a725e2b35cf7230d6677ffee\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000001ac7c4159f72b900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002e640d76638740f\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x5d\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0xd97b1de3619ed2c6beb3860147e30ca8a7dc9891\",\n      \"topics\": [\n        \"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n        \"0x000000000000000000000000d97b1de3619ed2c6beb3860147e30ca8a7dc9891\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000000015059f36c8ec0\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x5e\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0xd97b1de3619ed2c6beb3860147e30ca8a7dc9891\",\n      \"topics\": [\n        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n        \"0x000000000000000000000000735b14bb79463307aacbed86daf3322b1e6226ab\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000000015059f36c8ec0\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x5f\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0xd97b1de3619ed2c6beb3860147e30ca8a7dc9891\",\n      \"topics\": [\n        \"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n        \"0x000000000000000000000000d97b1de3619ed2c6beb3860147e30ca8a7dc9891\"\n      ],\n      \"data\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x60\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0xd97b1de3619ed2c6beb3860147e30ca8a7dc9891\",\n      \"topics\": [\n        \"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n        \"0x0000000000000000000000000000000000000000000000000000000000000000\"\n      ],\n      \"data\": \"0x00000000000000000000000000000000000000000000000002e4f07d72cbe54f\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x61\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0xd97b1de3619ed2c6beb3860147e30ca8a7dc9891\",\n      \"topics\": [\n        \"0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955\",\n        \"0x0000000000000000000000008e0f8e7e9e121403e72151d00f4937eacb2d9ef3\"\n      ],\n      \"data\": \"0x000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000002e4f07d72cbe54f00000000000000000000000000000000000000000000000000015059f36c8ec0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000005dabfdd153aaab4a970fd953dcfeee8bf6bb946e\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x62\",\n      \"removed\": false\n    },\n    {\n      \"address\": \"0x8e0f8e7e9e121403e72151d00f4937eacb2d9ef3\",\n      \"topics\": [\n        \"0x97eb75cc53ffa3f4560fc62e4912dda10ac56c3d12dbc48dc8c27d5ab225cf66\"\n      ],\n      \"data\": \"0x0000000000000000000000005f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf000000000000000000000000d97b1de3619ed2c6beb3860147e30ca8a7dc989100000000000000000000000000000000000000000000001b0d04202f47ec000000000000000000000000000000000000000000000000001ac7c4159f72b900000000000000000000000000005dabfdd153aaab4a970fd953dcfeee8bf6bb946e00000000000000000000000000000000000000000000000045400a8fd5330000\",\n      \"blockNumber\": \"0x17ef22\",\n      \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n      \"transactionIndex\": \"0x1f\",\n      \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n      \"logIndex\": \"0x63\",\n      \"removed\": false\n    }\n  ],\n  \"transactionHash\": \"0x87229bb05d67f42017a697b34ed52d95afc9f5e3285479e845fe088b4c77d8f0\",\n  \"contractAddress\": \"0x0000000000000000000000000000000000000000\",\n  \"gasUsed\": \"0x41c3c\",\n  \"blockHash\": \"0xf49e7039c7f1a81cd46de150980d92fa869cc0d2e2f1fe46aedc6400396137ff\",\n  \"blockNumber\": \"0x17ef22\",\n  \"transactionIndex\": \"0x1f\"\n}"
	err := json.Unmarshal([]byte(block), &receipt)
	require.NoError(t, err)
	return

}

// receiver is bc1qysd4sp9q8my59ul9wsf5rvs9p387hf8vfwatzu
func GetValidZRC20WithdrawToBTC(t *testing.T) (receipt ethtypes.Receipt) {
	block := "{\"type\":\"0x2\",\"root\":\"0x\",\"status\":\"0x1\",\"cumulativeGasUsed\":\"0x1f25ed\",\"logsBloom\":\"0x00000000000000000000000000020000000000000000000000000000000000020000000100000000000000000040000080000000000000000000000400200000200000000002000000000008000000000000000000000000000000000000000000000000020000000000000000800800000000000000000000000010000000000000000000000000000000000000000000000000000004000000000000000000020000000001000000000000000000000000000000000000000000000000010000000002000000000000000010000000000000000000000000000000000020000010000000000000000000000000000000000000040200000000000000000000\",\"logs\":[{\"address\":\"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x00000000000000000000000033ead83db0d0c682b05ead61e8d8f481bb1b4933\",\"0x000000000000000000000000735b14bb79463307aacbed86daf3322b1e6226ab\"],\"data\":\"0x0000000000000000000000000000000000000000000000000000000000003d84\",\"blockNumber\":\"0x1a00f3\",\"transactionHash\":\"0x9aaefece38fd2bd87077038a63fffb7c84cc8dd1ed01de134a8504a1f9a410c3\",\"transactionIndex\":\"0x8\",\"blockHash\":\"0x9517356f0b3877990590421266f02a4ff349b7476010ee34dd5f0dfc85c2684f\",\"logIndex\":\"0x28\",\"removed\":false},{\"address\":\"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\"topics\":[\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\"0x00000000000000000000000033ead83db0d0c682b05ead61e8d8f481bb1b4933\",\"0x00000000000000000000000013a0c5930c028511dc02665e7285134b6d11a5f4\"],\"data\":\"0x0000000000000000000000000000000000000000000000000000000000978c98\",\"blockNumber\":\"0x1a00f3\",\"transactionHash\":\"0x9aaefece38fd2bd87077038a63fffb7c84cc8dd1ed01de134a8504a1f9a410c3\",\"transactionIndex\":\"0x8\",\"blockHash\":\"0x9517356f0b3877990590421266f02a4ff349b7476010ee34dd5f0dfc85c2684f\",\"logIndex\":\"0x29\",\"removed\":false},{\"address\":\"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x00000000000000000000000033ead83db0d0c682b05ead61e8d8f481bb1b4933\",\"0x0000000000000000000000000000000000000000000000000000000000000000\"],\"data\":\"0x0000000000000000000000000000000000000000000000000000000000003039\",\"blockNumber\":\"0x1a00f3\",\"transactionHash\":\"0x9aaefece38fd2bd87077038a63fffb7c84cc8dd1ed01de134a8504a1f9a410c3\",\"transactionIndex\":\"0x8\",\"blockHash\":\"0x9517356f0b3877990590421266f02a4ff349b7476010ee34dd5f0dfc85c2684f\",\"logIndex\":\"0x2a\",\"removed\":false},{\"address\":\"0x13a0c5930c028511dc02665e7285134b6d11a5f4\",\"topics\":[\"0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955\",\"0x00000000000000000000000033ead83db0d0c682b05ead61e8d8f481bb1b4933\"],\"data\":\"0x000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000003d840000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a626331717973643473703971386d793539756c3977736635727673397033383768663876667761747a7500000000000000000000000000000000000000000000\",\"blockNumber\":\"0x1a00f3\",\"transactionHash\":\"0x9aaefece38fd2bd87077038a63fffb7c84cc8dd1ed01de134a8504a1f9a410c3\",\"transactionIndex\":\"0x8\",\"blockHash\":\"0x9517356f0b3877990590421266f02a4ff349b7476010ee34dd5f0dfc85c2684f\",\"logIndex\":\"0x2b\",\"removed\":false}],\"transactionHash\":\"0x9aaefece38fd2bd87077038a63fffb7c84cc8dd1ed01de134a8504a1f9a410c3\",\"contractAddress\":\"0x0000000000000000000000000000000000000000\",\"gasUsed\":\"0x12575\",\"blockHash\":\"0x9517356f0b3877990590421266f02a4ff349b7476010ee34dd5f0dfc85c2684f\",\"blockNumber\":\"0x1a00f3\",\"transactionIndex\":\"0x8\"}\n"
	err := json.Unmarshal([]byte(block), &receipt)
	require.NoError(t, err)
	return
}

func GetValidZetaSentDestinationExternal(t *testing.T) (receipt ethtypes.Receipt) {
	block := "{\"root\":\"0x\",\"status\":\"0x1\",\"cumulativeGasUsed\":\"0xd75f4f\",\"logsBloom\":\"0x00000000000000000000000000000000800800000000000000000000100000000000002000000100000000000000000000000000000000000000000000240000000000000000000000000008000000000800000000440000000000008080000000000000000000000000000000000000000000000000040000000010000000000000000000000000000000000000000200000001000000000000000040000000020000000000000000000000008200000000000000000000000000000000000000000002000000000000008000000000000000000000000000080002000041000010000000000000000000000000000000000000000000400000000000000000\",\"logs\":[{\"address\":\"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\"topics\":[\"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c\",\"0x000000000000000000000000f0a3f93ed1b126142e61423f9546bf1323ff82df\"],\"data\":\"0x000000000000000000000000000000000000000000000003cb71f51fc5580000\",\"blockNumber\":\"0x1bedc8\",\"transactionHash\":\"0x19d8a67a05998f1cb19fe731b96d817d5b186b62c9430c51679664959c952ef0\",\"transactionIndex\":\"0x5f\",\"blockHash\":\"0x198fdd1f4bc6b910db978602cb15bdb2bcc6fd960e9324e9b9675dc062133794\",\"logIndex\":\"0x13b\",\"removed\":false},{\"address\":\"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\"topics\":[\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\"0x000000000000000000000000f0a3f93ed1b126142e61423f9546bf1323ff82df\",\"0x000000000000000000000000239e96c8f17c85c30100ac26f635ea15f23e9c67\"],\"data\":\"0x000000000000000000000000000000000000000000000003cb71f51fc5580000\",\"blockNumber\":\"0x1bedc8\",\"transactionHash\":\"0x19d8a67a05998f1cb19fe731b96d817d5b186b62c9430c51679664959c952ef0\",\"transactionIndex\":\"0x5f\",\"blockHash\":\"0x198fdd1f4bc6b910db978602cb15bdb2bcc6fd960e9324e9b9675dc062133794\",\"logIndex\":\"0x13c\",\"removed\":false},{\"address\":\"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000f0a3f93ed1b126142e61423f9546bf1323ff82df\",\"0x000000000000000000000000239e96c8f17c85c30100ac26f635ea15f23e9c67\"],\"data\":\"0x000000000000000000000000000000000000000000000003cb71f51fc5580000\",\"blockNumber\":\"0x1bedc8\",\"transactionHash\":\"0x19d8a67a05998f1cb19fe731b96d817d5b186b62c9430c51679664959c952ef0\",\"transactionIndex\":\"0x5f\",\"blockHash\":\"0x198fdd1f4bc6b910db978602cb15bdb2bcc6fd960e9324e9b9675dc062133794\",\"logIndex\":\"0x13d\",\"removed\":false},{\"address\":\"0x5f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf\",\"topics\":[\"0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65\",\"0x000000000000000000000000239e96c8f17c85c30100ac26f635ea15f23e9c67\"],\"data\":\"0x000000000000000000000000000000000000000000000003cb71f51fc5580000\",\"blockNumber\":\"0x1bedc8\",\"transactionHash\":\"0x19d8a67a05998f1cb19fe731b96d817d5b186b62c9430c51679664959c952ef0\",\"transactionIndex\":\"0x5f\",\"blockHash\":\"0x198fdd1f4bc6b910db978602cb15bdb2bcc6fd960e9324e9b9675dc062133794\",\"logIndex\":\"0x13e\",\"removed\":false},{\"address\":\"0x239e96c8f17c85c30100ac26f635ea15f23e9c67\",\"topics\":[\"0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4\",\"0x000000000000000000000000f0a3f93ed1b126142e61423f9546bf1323ff82df\",\"0x0000000000000000000000000000000000000000000000000000000000000001\"],\"data\":\"0x00000000000000000000000060983881bdf302dcfa96603a58274d15d596620900000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000003cb71f51fc558000000000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000001460983881bdf302dcfa96603a58274d15d59662090000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000\",\"blockNumber\":\"0x1bedc8\",\"transactionHash\":\"0x19d8a67a05998f1cb19fe731b96d817d5b186b62c9430c51679664959c952ef0\",\"transactionIndex\":\"0x5f\",\"blockHash\":\"0x198fdd1f4bc6b910db978602cb15bdb2bcc6fd960e9324e9b9675dc062133794\",\"logIndex\":\"0x13f\",\"removed\":false}],\"transactionHash\":\"0x19d8a67a05998f1cb19fe731b96d817d5b186b62c9430c51679664959c952ef0\",\"contractAddress\":\"0x0000000000000000000000000000000000000000\",\"gasUsed\":\"0x2406d\",\"blockHash\":\"0x198fdd1f4bc6b910db978602cb15bdb2bcc6fd960e9324e9b9675dc062133794\",\"blockNumber\":\"0x1bedc8\",\"transactionIndex\":\"0x5f\"}\n"
	err := json.Unmarshal([]byte(block), &receipt)
	require.NoError(t, err)
	return
}
