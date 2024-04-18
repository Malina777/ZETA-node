package keeper

import (
	"math/big"
	"strings"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/pkg/chains"
	"github.com/zeta-chain/zetacore/pkg/coin"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

// hardcoded rate limiter flags
var rateLimitFlags = types.RateLimiterFlags{
	IsEnabled:       true,
	RateLimitWindow: 1200,                  // 1200 zeta blocks, 2 hours
	RateLimitInZeta: math.NewUint(2000000), // 2,000,000 ZETA
	Zrc20Rates: []*types.ZRC20Rate{
		// ETH
		{
			ChainId:        chains.GoerliLocalnetChain().ChainId,
			CoinType:       coin.CoinType_Gas,
			Asset:          "",
			ConversionRate: 2500,
		},
		// USDT
		{
			ChainId:        chains.GoerliLocalnetChain().ChainId,
			CoinType:       coin.CoinType_ERC20,
			Asset:          "0xbD1e64A22B9F92D9Ce81aA9B4b0fFacd80215564",
			ConversionRate: 0.8,
		},
		// BTC
		{
			ChainId:        chains.BtcRegtestChain().ChainId,
			CoinType:       coin.CoinType_ERC20,
			Asset:          "0x8f56682c2b8b2e3d4f6f7f7d6f3c01b3f6f6a7d6",
			ConversionRate: 50000,
		},
	},
}

// SetRatelimiterFlags set the rate limiter flags in the store
func (k Keeper) SetRatelimiterFlags(ctx sdk.Context, crosschainFlags types.RateLimiterFlags) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RateLimiterFlagsKey))
	b := k.cdc.MustMarshal(&crosschainFlags)
	store.Set([]byte{0}, b)
}

// GetRatelimiterFlags read the rate limiter flags from the store
func (k Keeper) GetRatelimiterFlags(_ sdk.Context) (val types.RateLimiterFlags, found bool) {
	return rateLimitFlags, true
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RateLimiterFlagsKey))

	// b := store.Get([]byte{0})
	// if b == nil {
	// 	return val, false
	// }

	// k.cdc.MustUnmarshal(b, &val)
	// return val, true
}

// GetRatelimiterRates returns two maps of foreign coins and their rates
// The 1st map: foreign chain id -> gas coin rate
// The 2nd map: foreign erc20 asset -> erc20 coin rate
func (k Keeper) GetRatelimiterRates(ctx sdk.Context) (map[int64]*big.Float, map[string]*big.Float) {
	rateLimitFlags, _ := k.GetRatelimiterFlags(ctx)

	// the result maps
	gasCoinRates := make(map[int64]*big.Float)
	erc20CoinRates := make(map[string]*big.Float)

	// loop through the rate limiter flags to get the rate
	for _, zrc20Rate := range rateLimitFlags.Zrc20Rates {
		rate := big.NewFloat(zrc20Rate.ConversionRate)
		switch zrc20Rate.CoinType {
		case coin.CoinType_Gas:
			gasCoinRates[zrc20Rate.ChainId] = rate
		case coin.CoinType_ERC20:
			erc20CoinRates[strings.ToLower(zrc20Rate.Asset)] = rate
		}
	}
	return gasCoinRates, erc20CoinRates
}
