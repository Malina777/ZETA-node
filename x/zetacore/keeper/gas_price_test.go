package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/zeta-chain/zetacore/x/zetacore/types"
)

func createNGasPrice(keeper *Keeper, ctx sdk.Context, n int) []types.GasPrice {
	items := make([]types.GasPrice, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetGasPrice(ctx, items[i])
	}
	return items
}

func TestGasPriceGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNGasPrice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGasPrice(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestGasPriceRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNGasPrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGasPrice(ctx, item.Index)
		_, found := keeper.GetGasPrice(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestGasPriceGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNGasPrice(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllGasPrice(ctx))
}
