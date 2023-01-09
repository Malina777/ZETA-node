package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	types2 "github.com/zeta-chain/zetacore/x/observer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
)

func (k Keeper) ConvertGasToZeta(context context.Context, request *types.QueryConvertGasToZetaRequest) (*types.QueryConvertGasToZetaResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	chainName := types2.ParseStringToObserverChain(request.Chain)
	chain, _ := k.zetaObserverKeeper.GetChainFromChainName(ctx, chainName)
	medianGasPrice, isFound := k.GetMedianGasPriceInUint(ctx, chain.ChainId)
	if !isFound {
		return nil, status.Error(codes.InvalidArgument, "invalid request: param chain")
	}

	gasLimit := sdk.NewUintFromString(request.GasLimit)
	outTxGasFee := medianGasPrice.Mul(gasLimit)
	zrc20, err := k.fungibleKeeper.QuerySystemContractGasCoinZRC4(ctx, big.NewInt(chain.ChainId))
	if err != nil {
		return nil, status.Error(codes.NotFound, "zrc20 not found")
	}
	outTxGasFeeInZeta, err := k.fungibleKeeper.QueryUniswapv2RouterGetAmountsIn(ctx, outTxGasFee.BigInt(), zrc20)
	if err != nil {
		return nil, status.Error(codes.Internal, "zQueryUniswapv2RouterGetAmountsIn failed")
	}
	return &types.QueryConvertGasToZetaResponse{
		OutboundGasInZeta: outTxGasFeeInZeta.String(),
		ProtocolFeeInZeta: types.GetProtocolFee().String(),
		ZetaBlockHeight:   uint64(ctx.BlockHeight()),
	}, nil
}

func (k Keeper) ProtocolFee(context context.Context, req *types.QueryMessagePassingProtocolFeeRequest) (*types.QueryMessagePassingProtocolFeeResponse, error) {
	return &types.QueryMessagePassingProtocolFeeResponse{
		FeeInZeta: types.GetProtocolFee().String(),
	}, nil
}
