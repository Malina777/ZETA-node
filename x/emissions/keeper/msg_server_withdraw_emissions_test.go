package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/cmd/zetacored/config"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/emissions/keeper"
	"github.com/zeta-chain/zetacore/x/emissions/types"
)

func TestMsgServer_WithdrawEmission(t *testing.T) {
	t.Run("successfully withdraw emissions", func(t *testing.T) {
		k, ctx, sk, _ := keepertest.EmissionsKeeper(t)

		msgServer := keeper.NewMsgServerImpl(*k)
		withdrawableEmission := sample.WithdrawableEmissions(t)
		k.SetWithdrawableEmission(ctx, withdrawableEmission)
		err := sk.BankKeeper.MintCoins(ctx, types.UndistributedObserverRewardsPool, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, withdrawableEmission.Amount)))
		require.NoError(t, err)

		_, err = msgServer.WithdrawEmission(ctx, &types.MsgWithdrawEmission{
			Creator: withdrawableEmission.Address,
			Amount:  withdrawableEmission.Amount,
		})
		require.NoError(t, err)

		balance := k.GetBankKeeper().GetBalance(ctx, sdk.MustAccAddressFromBech32(withdrawableEmission.Address), config.BaseDenom).Amount.String()
		require.Equal(t, withdrawableEmission.Amount.String(), balance)
	})

	t.Run("unable to create withdraw emissions with invalid address", func(t *testing.T) {
		k, ctx, sk, _ := keepertest.EmissionsKeeper(t)

		msgServer := keeper.NewMsgServerImpl(*k)
		withdrawableEmission := sample.WithdrawableEmissions(t)
		k.SetWithdrawableEmission(ctx, withdrawableEmission)
		err := sk.BankKeeper.MintCoins(ctx, types.UndistributedObserverRewardsPool, sdk.NewCoins(sdk.NewCoin(config.BaseDenom, withdrawableEmission.Amount)))
		require.NoError(t, err)

		_, err = msgServer.WithdrawEmission(ctx, &types.MsgWithdrawEmission{
			Creator: "invalid_address",
			Amount:  withdrawableEmission.Amount,
		})
		require.ErrorIs(t, err, types.ErrInvalidAddress)
	})

	t.Run("unable to create withdraw emissions if undistributed rewards pool does not have enough balance", func(t *testing.T) {
		k, ctx, _, _ := keepertest.EmissionsKeeper(t)

		msgServer := keeper.NewMsgServerImpl(*k)
		withdrawableEmission := sample.WithdrawableEmissions(t)
		k.SetWithdrawableEmission(ctx, withdrawableEmission)

		_, err := msgServer.WithdrawEmission(ctx, &types.MsgWithdrawEmission{
			Creator: withdrawableEmission.Address,
			Amount:  withdrawableEmission.Amount,
		})
		require.ErrorIs(t, err, types.ErrRewardsPoolDoesNotHaveEnoughBalance)
	})

	t.Run("unable to create withdraw emissions with invalid amount", func(t *testing.T) {
		k, ctx, _, _ := keepertest.EmissionsKeeper(t)
		msgServer := keeper.NewMsgServerImpl(*k)
		withdrawableEmission := sample.WithdrawableEmissions(t)
		k.SetWithdrawableEmission(ctx, withdrawableEmission)
		_, err := msgServer.WithdrawEmission(ctx, &types.MsgWithdrawEmission{
			Creator: withdrawableEmission.Address,
			Amount:  sdkmath.NewInt(-1),
		})
		require.ErrorIs(t, err, types.ErrUnableToWithdrawEmissions)
	})

}