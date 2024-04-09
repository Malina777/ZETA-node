package keeper_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/x/evm/statedb"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/testutil/contracts"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/fungible/types"
)

func TestKeeper_ZevmOnReceive(t *testing.T) {
	t.Run("successfully call ZevmOnReceive on connector contract ", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		deploySystemContracts(t, ctx, k, sdkk.EvmKeeper)
		dAppContract, err := k.DeployContract(ctx, contracts.DappMetaData)
		require.NoError(t, err)
		assertContractDeployment(t, sdkk.EvmKeeper, ctx, dAppContract)

		zetaTxSender := sample.EthAddress().Bytes()
		senderChainID := big.NewInt(1)
		zetaTxReceiver := dAppContract
		amount := big.NewInt(45)
		data := []byte("message")
		cctxIndexBytes := [32]byte{}

		_, err = k.ZevmOnReceive(ctx, zetaTxSender, zetaTxReceiver, senderChainID, amount, data, cctxIndexBytes)
		require.NoError(t, err)

		dappAbi, err := contracts.DappMetaData.GetAbi()
		require.NoError(t, err)
		res, err := k.CallEVM(
			ctx,
			*dappAbi,
			types.ModuleAddressEVM,
			dAppContract,
			big.NewInt(0),
			nil,
			false,
			false,
			"zetaTxSenderAddress",
		)
		require.NoError(t, err)
		unpacked, err := dappAbi.Unpack("zetaTxSenderAddress", res.Ret)
		require.NoError(t, err)
		require.NotZero(t, len(unpacked))
		valSenderAddress, ok := unpacked[0].([]byte)
		require.True(t, ok)
		require.Equal(t, zetaTxSender, valSenderAddress)
	})

	t.Run("fail to call ZevmOnReceive if account not found for receiver address", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		deploySystemContracts(t, ctx, k, sdkk.EvmKeeper)

		_, err := k.ZevmOnReceive(ctx, sample.EthAddress().Bytes(),
			sample.EthAddress(),
			big.NewInt(1),
			big.NewInt(45),
			[]byte("message"),
			[32]byte{})
		require.ErrorIs(t, err, types.ErrAccountNotFound)
	})

	t.Run("fail to call ZevmOnReceive if account is not a contract", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		deploySystemContracts(t, ctx, k, sdkk.EvmKeeper)

		zetaTxReceiver := sample.EthAddress()
		err := sdkk.EvmKeeper.SetAccount(ctx, zetaTxReceiver, statedb.Account{
			Nonce:    0,
			Balance:  big.NewInt(100),
			CodeHash: crypto.Keccak256(nil),
		})
		require.NoError(t, err)

		_, err = k.ZevmOnReceive(ctx, sample.EthAddress().Bytes(),
			zetaTxReceiver,
			big.NewInt(1),
			big.NewInt(45),
			[]byte("message"),
			[32]byte{})
		require.ErrorIs(t, err, types.ErrCallNonContract)
	})

	t.Run("fail to call ZevmOnReceive if CallOnReceiveZevmConnector fails", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		dAppContract, err := k.DeployContract(ctx, contracts.DappMetaData)
		require.NoError(t, err)
		assertContractDeployment(t, sdkk.EvmKeeper, ctx, dAppContract)

		zetaTxSender := sample.EthAddress().Bytes()
		senderChainID := big.NewInt(1)
		zetaTxReceiver := dAppContract
		amount := big.NewInt(45)
		data := []byte("message")
		cctxIndexBytes := [32]byte{}

		_, err = k.ZevmOnReceive(ctx, zetaTxSender, zetaTxReceiver, senderChainID, amount, data, cctxIndexBytes)
		require.ErrorIs(t, err, types.ErrContractNotFound)
		require.ErrorContains(t, err, "GetSystemContract address not found")
	})
}

func TestKeeper_ZevmOnRevert(t *testing.T) {
	t.Run("successfully call ZevmOnRevert on connector contract ", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		deploySystemContracts(t, ctx, k, sdkk.EvmKeeper)
		dAppContract, err := k.DeployContract(ctx, contracts.DappMetaData)
		require.NoError(t, err)
		assertContractDeployment(t, sdkk.EvmKeeper, ctx, dAppContract)

		zetaTxSender := dAppContract
		senderChainID := big.NewInt(1)
		destinationChainID := big.NewInt(2)
		zetaTxReceiver := sample.EthAddress().Bytes()
		amount := big.NewInt(45)
		data := []byte("message")
		cctxIndexBytes := [32]byte{}

		_, err = k.ZevmOnRevert(ctx, zetaTxSender, zetaTxReceiver, senderChainID, destinationChainID, amount, data, cctxIndexBytes)
		require.NoError(t, err)

		dappAbi, err := contracts.DappMetaData.GetAbi()
		require.NoError(t, err)
		res, err := k.CallEVM(
			ctx,
			*dappAbi,
			types.ModuleAddressEVM,
			dAppContract,
			big.NewInt(0),
			nil,
			false,
			false,
			"zetaTxSenderAddress",
		)
		require.NoError(t, err)
		unpacked, err := dappAbi.Unpack("zetaTxSenderAddress", res.Ret)
		require.NoError(t, err)
		require.NotZero(t, len(unpacked))
		valSenderAddress, ok := unpacked[0].([]byte)
		require.True(t, ok)
		require.Equal(t, zetaTxSender.Bytes(), valSenderAddress)
	})

	t.Run("fail to call ZevmOnRevert if account is not a contract", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		deploySystemContracts(t, ctx, k, sdkk.EvmKeeper)

		zetaTxSender := sample.EthAddress()
		err := sdkk.EvmKeeper.SetAccount(ctx, zetaTxSender, statedb.Account{
			Nonce:    0,
			Balance:  big.NewInt(100),
			CodeHash: crypto.Keccak256(nil),
		})
		require.NoError(t, err)

		_, err = k.ZevmOnRevert(ctx, zetaTxSender, sample.EthAddress().Bytes(),
			big.NewInt(1),
			big.NewInt(2),
			big.NewInt(45),
			[]byte("message"),
			[32]byte{})
		require.ErrorIs(t, err, types.ErrCallNonContract)
	})

	t.Run("fail to call ZevmOnRevert if CallOnRevertZevmConnector fails", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		dAppContract, err := k.DeployContract(ctx, contracts.DappMetaData)
		require.NoError(t, err)
		assertContractDeployment(t, sdkk.EvmKeeper, ctx, dAppContract)

		zetaTxSender := dAppContract
		senderChainID := big.NewInt(1)
		destinationChainID := big.NewInt(2)
		zetaTxReceiver := sample.EthAddress().Bytes()
		amount := big.NewInt(45)
		data := []byte("message")
		cctxIndexBytes := [32]byte{}

		_, err = k.ZevmOnRevert(ctx, zetaTxSender, zetaTxReceiver, senderChainID, destinationChainID, amount, data, cctxIndexBytes)
		require.ErrorIs(t, err, types.ErrContractNotFound)
		require.ErrorContains(t, err, "GetSystemContract address not found")
	})

	t.Run("fail to call ZevmOnRevert if account not found for sender address", func(t *testing.T) {
		k, ctx, sdkk, _ := keepertest.FungibleKeeper(t)
		_ = k.GetAuthKeeper().GetModuleAccount(ctx, types.ModuleName)

		deploySystemContracts(t, ctx, k, sdkk.EvmKeeper)

		_, err := k.ZevmOnRevert(ctx, sample.EthAddress(),
			sample.EthAddress().Bytes(),
			big.NewInt(1),
			big.NewInt(2),
			big.NewInt(45),
			[]byte("message"),
			[32]byte{})
		require.ErrorIs(t, err, types.ErrAccountNotFound)
	})
}