package keeper

import (
	"context"
	"sort"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/pkg/coin"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RateLimiterInput collects the input data for the rate limiter
func (k Keeper) RateLimiterInput(c context.Context, req *types.QueryRateLimiterInputRequest) (res *types.QueryRateLimiterInputResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.Window <= 0 {
		return nil, status.Error(codes.InvalidArgument, "window must be positive")
	}

	// use default MaxPendingCctxs if not specified or too high
	limit := req.Limit
	if limit == 0 || limit > MaxPendingCctxs {
		limit = MaxPendingCctxs
	}
	ctx := sdk.UnwrapSDKContext(c)

	// get current height and tss
	height := ctx.BlockHeight()
	if height <= 0 {
		return nil, status.Error(codes.OutOfRange, "height out of range")
	}
	tss, found := k.zetaObserverKeeper.GetTSS(ctx)
	if !found {
		return nil, observertypes.ErrTssNotFound
	}

	// calculate the rate limiter sliding window left boundary (inclusive)
	leftWindowBoundary := height - req.Window + 1
	if leftWindowBoundary < 1 {
		leftWindowBoundary = 1
	}

	// the `limit` of pending result is reached or not
	maxCCTXsReached := func(cctxs []*types.CrossChainTx) bool {
		// #nosec G701 len always positive
		return uint32(len(cctxs)) > limit
	}

	// if a cctx falls within the rate limiter window
	isCctxInWindow := func(cctx *types.CrossChainTx) bool {
		// #nosec G701 checked positive
		return cctx.InboundTxParams.InboundTxObservedExternalHeight >= uint64(leftWindowBoundary)
	}

	// it is a past cctx if its nonce < `nonceLow`,
	isPastCctx := func(cctx *types.CrossChainTx, nonceLow int64) bool {
		// #nosec G701 always positive
		return cctx.GetCurrentOutTxParam().OutboundTxTssNonce < uint64(nonceLow)
	}

	// get foreign chains and conversion rates of foreign coins
	chains := k.zetaObserverKeeper.GetSupportedForeignChains(ctx)
	assetRates := k.GetRateLimiterAssetRateList(ctx)
	gasAssetRateMap, erc20AssetRateMap := BuildAssetRateMapFromList(assetRates)

	// query pending nonces of each foreign chain and get the lowest height of the pending cctxs
	lowestPendingCctxHeight := int64(0)
	pendingNoncesMap := make(map[int64]observertypes.PendingNonces)
	for _, chain := range chains {
		pendingNonces, found := k.GetObserverKeeper().GetPendingNonces(ctx, tss.TssPubkey, chain.ChainId)
		if !found {
			return nil, status.Error(codes.Internal, "pending nonces not found")
		}
		pendingNoncesMap[chain.ChainId] = pendingNonces

		// update lowest pending cctx height
		if pendingNonces.NonceLow < pendingNonces.NonceHigh {
			cctx, err := getCctxByChainIDAndNonce(k, ctx, tss.TssPubkey, chain.ChainId, pendingNonces.NonceLow)
			if err != nil {
				return nil, err
			}
			// #nosec G701 len always in range
			cctxHeight := int64(cctx.InboundTxParams.InboundTxObservedExternalHeight)
			if lowestPendingCctxHeight == 0 || cctxHeight < lowestPendingCctxHeight {
				lowestPendingCctxHeight = cctxHeight
			}
		}
	}

	// define a few variables to be used in the query loops
	totalPending := uint64(0)
	pastCctxsValue := sdk.NewInt(0)
	pendingCctxsValue := sdk.NewInt(0)
	cctxsMissed := make([]*types.CrossChainTx, 0)
	cctxsPending := make([]*types.CrossChainTx, 0)

	// query backwards for pending cctxs of each foreign chain
	for _, chain := range chains {
		// we should at least query 1000 prior to find any pending cctx that we might have missed
		// this logic is needed because a confirmation of higher nonce will automatically update the p.NonceLow
		// therefore might mask some lower nonce cctx that is still pending.
		pendingNonces := pendingNoncesMap[chain.ChainId]
		startNonce := pendingNonces.NonceHigh - 1
		endNonce := pendingNonces.NonceLow - MaxLookbackNonce
		if endNonce < 0 {
			endNonce = 0
		}

		// go all the way back to the left window boundary or `NonceLow - 1000`, depending on which on arrives first
		for nonce := startNonce; nonce >= 0; nonce-- {
			cctx, err := getCctxByChainIDAndNonce(k, ctx, tss.TssPubkey, chain.ChainId, nonce)
			if err != nil {
				return nil, err
			}
			inWindow := isCctxInWindow(cctx)
			isPast := isPastCctx(cctx, pendingNonces.NonceLow)

			// we should at least go backwards by 1000 nonces to pick up missed pending cctxs
			// we might go even further back if the endNonce hasn't hit the left window boundary yet
			if nonce < endNonce && !inWindow {
				break
			}

			// sum up past cctxs' value within window
			if inWindow && isPast {
				pastCctxsValue = pastCctxsValue.Add(ConvertCctxValue(chain.ChainId, cctx, gasAssetRateMap, erc20AssetRateMap))
			}

			// add cctx to corresponding list
			if IsPending(cctx) {
				totalPending++
				if isPast {
					cctxsMissed = append(cctxsMissed, cctx)
				} else {
					cctxsPending = append(cctxsPending, cctx)
					// sum up non-past pending cctxs' value
					pendingCctxsValue = pendingCctxsValue.Add(ConvertCctxValue(chain.ChainId, cctx, gasAssetRateMap, erc20AssetRateMap))
				}
			}
		}
	}

	// sort the missed cctxs order by height (can sort by other criteria, for unit testability)
	SortCctxsByHeightAndChainID(cctxsMissed)

	// sort the pending cctxs order by height (first come first serve)
	SortCctxsByHeightAndChainID(cctxsPending)

	// we take all the missed cctxs (won't be a lot) for simplicity of the query, but we only take a `limit` number of pending cctxs
	if maxCCTXsReached(cctxsPending) {
		cctxsPending = cctxsPending[:limit]
	}

	return &types.QueryRateLimiterInputResponse{
		Height:                  height,
		CctxsMissed:             cctxsMissed,
		CctxsPending:            cctxsPending,
		TotalPending:            totalPending,
		PastCctxsValue:          pastCctxsValue.String(),
		PendingCctxsValue:       pendingCctxsValue.String(),
		LowestPendingCctxHeight: lowestPendingCctxHeight,
	}, nil
}

// ListPendingCctxWithinRateLimit returns a list of pending cctxs that do not exceed the outbound rate limit
// a limit for the number of cctxs to return can be specified or the default is MaxPendingCctxs
func (k Keeper) ListPendingCctxWithinRateLimit(c context.Context, req *types.QueryListPendingCctxWithinRateLimitRequest) (res *types.QueryListPendingCctxWithinRateLimitResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// use default MaxPendingCctxs if not specified or too high
	limit := req.Limit
	if limit == 0 || limit > MaxPendingCctxs {
		limit = MaxPendingCctxs
	}
	ctx := sdk.UnwrapSDKContext(c)

	// define a few variables to be used in the query loops
	limitExceeded := false
	totalPending := uint64(0)
	totalWithdrawInAzeta := sdkmath.NewInt(0)
	cctxs := make([]*types.CrossChainTx, 0)
	chains := k.zetaObserverKeeper.GetSupportedForeignChains(ctx)

	// check rate limit flags to decide if we should apply rate limit
	applyLimit := true
	rateLimitFlags, found := k.GetRateLimiterFlags(ctx)
	if !found || !rateLimitFlags.Enabled {
		applyLimit = false
	}
	if rateLimitFlags.Rate.IsNil() || rateLimitFlags.Rate.IsZero() {
		applyLimit = false
	}

	// fallback to non-rate-limited query if rate limiter is disabled
	if !applyLimit {
		for _, chain := range chains {
			resp, err := k.ListPendingCctx(ctx, &types.QueryListPendingCctxRequest{ChainId: chain.ChainId, Limit: limit})
			if err == nil {
				cctxs = append(cctxs, resp.CrossChainTx...)
				totalPending += resp.TotalPending
			}
		}
		return &types.QueryListPendingCctxWithinRateLimitResponse{
			CrossChainTx:      cctxs,
			TotalPending:      totalPending,
			RateLimitExceeded: false,
		}, nil
	}

	// get current height and tss
	height := ctx.BlockHeight()
	if height <= 0 {
		return nil, status.Error(codes.OutOfRange, "height out of range")
	}
	tss, found := k.zetaObserverKeeper.GetTSS(ctx)
	if !found {
		return nil, observertypes.ErrTssNotFound
	}

	// calculate the rate limiter sliding window left boundary (inclusive)
	leftWindowBoundary := height - rateLimitFlags.Window + 1
	if leftWindowBoundary < 0 {
		leftWindowBoundary = 0
	}

	// get the conversion rates for all foreign coins
	var gasAssetRateMap map[int64]*types.AssetRate
	var erc20AssetRateMap map[int64]map[string]*types.AssetRate
	var blockLimitInAzeta sdkmath.Int
	var windowLimitInAzeta sdkmath.Int
	if applyLimit {
		assetRates := k.GetRateLimiterAssetRateList(ctx)
		gasAssetRateMap, erc20AssetRateMap = BuildAssetRateMapFromList(assetRates)

		// initiate block limit and window limit in azeta
		blockLimitInAzeta = sdkmath.NewIntFromBigInt(rateLimitFlags.Rate.BigInt())
		windowLimitInAzeta = blockLimitInAzeta.Mul(sdkmath.NewInt(rateLimitFlags.Window))
	}

	// the criteria to stop adding cctxs to the rpc response
	maxCCTXsReached := func(cctxs []*types.CrossChainTx) bool {
		// #nosec G701 len always positive
		return uint32(len(cctxs)) >= limit
	}

	// if a cctx falls within the rate limiter window
	isCctxInWindow := func(cctx *types.CrossChainTx) bool {
		// #nosec G701 checked positive
		return cctx.InboundTxParams.InboundTxObservedExternalHeight >= uint64(leftWindowBoundary)
	}

	// query pending nonces for each foreign chain and get the lowest height of the pending cctxs
	lowestPendingCctxHeight := int64(0)
	pendingNoncesMap := make(map[int64]observertypes.PendingNonces)
	for _, chain := range chains {
		pendingNonces, found := k.GetObserverKeeper().GetPendingNonces(ctx, tss.TssPubkey, chain.ChainId)
		if !found {
			return nil, status.Error(codes.Internal, "pending nonces not found")
		}
		pendingNoncesMap[chain.ChainId] = pendingNonces

		// insert pending nonces and update lowest height
		if pendingNonces.NonceLow < pendingNonces.NonceHigh {
			cctx, err := getCctxByChainIDAndNonce(k, ctx, tss.TssPubkey, chain.ChainId, pendingNonces.NonceLow)
			if err != nil {
				return nil, err
			}
			// #nosec G701 len always in range
			cctxHeight := int64(cctx.InboundTxParams.InboundTxObservedExternalHeight)
			if lowestPendingCctxHeight == 0 || cctxHeight < lowestPendingCctxHeight {
				lowestPendingCctxHeight = cctxHeight
			}
		}
	}

	// invariant: for period of time >= `rateLimitFlags.Window`, the zetaclient-side average withdraw rate should be <= `blockLimitInZeta`
	// otherwise, this query should return empty result and wait for the average rate to drop below `blockLimitInZeta`
	withdrawWindow := rateLimitFlags.Window
	withdrawLimitInAzeta := windowLimitInAzeta
	if lowestPendingCctxHeight != 0 {
		// `pendingCctxWindow` is the width of [lowestPendingCctxHeight, height] window
		// if the window can be wider than `rateLimitFlags.Window`, we should adjust the total withdraw limit proportionally
		pendingCctxWindow := height - lowestPendingCctxHeight + 1
		if pendingCctxWindow > rateLimitFlags.Window {
			withdrawWindow = pendingCctxWindow
			withdrawLimitInAzeta = blockLimitInAzeta.Mul(sdk.NewInt(pendingCctxWindow))
		}
	}

	// query backwards for potential missed pending cctxs for each foreign chain
	for _, chain := range chains {
		// we should at least query 1000 prior to find any pending cctx that we might have missed
		// this logic is needed because a confirmation of higher nonce will automatically update the p.NonceLow
		// therefore might mask some lower nonce cctx that is still pending.
		pendingNonces := pendingNoncesMap[chain.ChainId]
		startNonce := pendingNonces.NonceLow - 1
		endNonce := pendingNonces.NonceLow - MaxLookbackNonce
		if endNonce < 0 {
			endNonce = 0
		}

		// query cctx by nonce backwards to the left boundary of the rate limit sliding window
		for nonce := startNonce; nonce >= 0; nonce-- {
			cctx, err := getCctxByChainIDAndNonce(k, ctx, tss.TssPubkey, chain.ChainId, nonce)
			if err != nil {
				return nil, err
			}
			inWindow := isCctxInWindow(cctx)

			// we should at least go backwards by 1000 nonces to pick up missed pending cctxs
			// we might go even further back if rate limiter is enabled and the endNonce hasn't hit the left window boundary yet
			// stop at the left window boundary if the `endNonce` hasn't hit it yet
			if nonce < endNonce && !inWindow {
				break
			}
			// skip the cctx if rate limit is exceeded but still accumulate the total withdraw value
			if inWindow && rateLimitExceeded(chain.ChainId, cctx, gasAssetRateMap, erc20AssetRateMap, &totalWithdrawInAzeta, withdrawLimitInAzeta) {
				limitExceeded = true
				continue
			}

			// only take a `limit` number of pending cctxs as result but still count the total pending cctxs
			if IsPending(cctx) {
				totalPending++
				if !maxCCTXsReached(cctxs) {
					cctxs = append(cctxs, cctx)
				}
			}
		}
	}

	// remember the number of missed pending cctxs
	missedPending := len(cctxs)

	// query forwards for pending cctxs for each foreign chain
	for _, chain := range chains {
		pendingNonces := pendingNoncesMap[chain.ChainId]

		// #nosec G701 always in range
		totalPending += uint64(pendingNonces.NonceHigh - pendingNonces.NonceLow)

		// query the pending cctxs in range [NonceLow, NonceHigh)
		for nonce := pendingNonces.NonceLow; nonce < pendingNonces.NonceHigh; nonce++ {
			cctx, err := getCctxByChainIDAndNonce(k, ctx, tss.TssPubkey, chain.ChainId, nonce)
			if err != nil {
				return nil, err
			}

			// skip the cctx if rate limit is exceeded but still accumulate the total withdraw value
			if rateLimitExceeded(chain.ChainId, cctx, gasAssetRateMap, erc20AssetRateMap, &totalWithdrawInAzeta, withdrawLimitInAzeta) {
				limitExceeded = true
				continue
			}
			// only take a `limit` number of pending cctxs as result
			if maxCCTXsReached(cctxs) {
				continue
			}
			cctxs = append(cctxs, cctx)
		}
	}

	// if the rate limit is exceeded, only return the missed pending cctxs
	if limitExceeded {
		cctxs = cctxs[:missedPending]
	}

	// sort the cctxs by chain ID and nonce (lower nonce holds higher priority for scheduling)
	sort.Slice(cctxs, func(i, j int) bool {
		if cctxs[i].GetCurrentOutTxParam().ReceiverChainId == cctxs[j].GetCurrentOutTxParam().ReceiverChainId {
			return cctxs[i].GetCurrentOutTxParam().OutboundTxTssNonce < cctxs[j].GetCurrentOutTxParam().OutboundTxTssNonce
		}
		return cctxs[i].GetCurrentOutTxParam().ReceiverChainId < cctxs[j].GetCurrentOutTxParam().ReceiverChainId
	})

	return &types.QueryListPendingCctxWithinRateLimitResponse{
		CrossChainTx:          cctxs,
		TotalPending:          totalPending,
		CurrentWithdrawWindow: withdrawWindow,
		CurrentWithdrawRate:   totalWithdrawInAzeta.Quo(sdk.NewInt(withdrawWindow)).String(),
		RateLimitExceeded:     limitExceeded,
	}, nil
}

// ConvertCctxValue converts the value of the cctx to azeta using given conversion rates
func ConvertCctxValue(
	chainID int64,
	cctx *types.CrossChainTx,
	gasAssetRateMap map[int64]*types.AssetRate,
	erc20AssetRateMap map[int64]map[string]*types.AssetRate,
) sdkmath.Int {
	var rate sdk.Dec
	var decimals uint32
	switch cctx.InboundTxParams.CoinType {
	case coin.CoinType_Zeta:
		// no conversion needed for ZETA
		return sdk.NewIntFromBigInt(cctx.GetCurrentOutTxParam().Amount.BigInt())
	case coin.CoinType_Gas:
		assetRate, found := gasAssetRateMap[chainID]
		if !found {
			// skip if no rate found for gas coin on this chainID
			return sdk.NewInt(0)
		}
		rate = assetRate.Rate
		decimals = assetRate.Decimals
	case coin.CoinType_ERC20:
		// get the ERC20 coin rate
		_, found := erc20AssetRateMap[chainID]
		if !found {
			// skip if no rate found for this chainID
			return sdk.NewInt(0)
		}
		assetRate := erc20AssetRateMap[chainID][strings.ToLower(cctx.InboundTxParams.Asset)]
		rate = assetRate.Rate
		decimals = assetRate.Decimals
	default:
		// skip CoinType_Cmd
		return sdk.NewInt(0)
	}
	// should not happen, return 0 to skip if it happens
	if rate.IsNil() || rate.LTE(sdk.NewDec(0)) {
		return sdkmath.NewInt(0)
	}

	// the whole coin amounts of zeta and zrc20
	// given decimals = 6, the amount will be 10^6 = 1000000
	oneZeta := coin.AzetaPerZeta()
	oneZrc20 := sdk.NewDec(10).Power(uint64(decimals))

	// convert cctx asset amount into azeta amount
	// given amountCctx = 2000000, rate = 0.8, decimals = 6
	// amountCctxDec: 2000000 * 0.8 = 1600000.0
	// amountAzetaDec: 1600000.0 * 10e18 / 10e6 = 1600000000000000000.0
	amountCctxDec := sdk.NewDecFromBigInt(cctx.GetCurrentOutTxParam().Amount.BigInt())
	amountAzetaDec := amountCctxDec.Mul(rate).Mul(oneZeta).Quo(oneZrc20)
	return amountAzetaDec.TruncateInt()
}

// rateLimitExceeded accumulates the cctx value and then checks if the rate limit is exceeded
// returns true if the rate limit is exceeded
func rateLimitExceeded(
	chainID int64,
	cctx *types.CrossChainTx,
	gasAssetRateMap map[int64]*types.AssetRate,
	erc20AssetRateMap map[int64]map[string]*types.AssetRate,
	currentCctxValue *sdkmath.Int,
	withdrawLimitInAzeta sdkmath.Int,
) bool {
	amountZeta := ConvertCctxValue(chainID, cctx, gasAssetRateMap, erc20AssetRateMap)
	*currentCctxValue = currentCctxValue.Add(amountZeta)
	return currentCctxValue.GT(withdrawLimitInAzeta)
}
