package keeper

import (
	"context"
	"cosmossdk.io/math"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//func (k Keeper) CctxChangePrefixStore(ctx sdk.Context, send types.CrossChainTx, oldStatus types.CctxStatus) {
//	// Defensive Programming :Remove first set later
//	_, found := k.GetCrossChainTx(ctx, send.Index, oldStatus)
//	if found {
//		k.RemoveCrossChainTx(ctx, send.Index, oldStatus)
//	}
//	k.SetCrossChainTx(ctx, send)
//}

// SetCrossChainTx set a specific send in the store from its index
func (k Keeper) SetCrossChainTx(ctx sdk.Context, send types.CrossChainTx) {

	p := types.KeyPrefix(fmt.Sprintf("%s", types.SendKey))
	store := prefix.NewStore(ctx.KVStore(k.storeKey), p)
	b := k.cdc.MustMarshal(&send)
	store.Set(types.KeyPrefix(send.Index), b)

	// set mapping inTxHash -> cctxIndex
	k.SetInTxHashToCctx(ctx, types.InTxHashToCctx{
		InTxHash:  send.InboundTxParams.InboundTxObservedHash,
		CctxIndex: send.Index,
	})

	tss, found := k.GetTSS(ctx)
	if !found {
		return
	}
	// set mapping nonce => cctxIndex
	if send.CctxStatus.Status == types.CctxStatus_PendingOutbound || send.CctxStatus.Status == types.CctxStatus_PendingRevert {
		k.SetNonceToCctx(ctx, types.NonceToCctx{
			ChainId:   send.GetCurrentOutTxParam().ReceiverChainId,
			Nonce:     int64(send.GetCurrentOutTxParam().OutboundTxTssNonce),
			CctxIndex: send.Index,
			Tss:       tss.TssPubkey,
		})
	}
}

// GetCrossChainTx returns a send from its index
func (k Keeper) GetCrossChainTx(ctx sdk.Context, index string) (val types.CrossChainTx, found bool) {
	p := types.KeyPrefix(fmt.Sprintf("%s", types.SendKey))
	store := prefix.NewStore(ctx.KVStore(k.storeKey), p)

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCrossChainTx removes a send from the store
func (k Keeper) RemoveCrossChainTx(ctx sdk.Context, index string) {
	p := types.KeyPrefix(fmt.Sprintf("%s", types.SendKey))
	store := prefix.NewStore(ctx.KVStore(k.storeKey), p)
	store.Delete(types.KeyPrefix(index))
}

// Queries

func (k Keeper) CctxAll(c context.Context, req *types.QueryAllCctxRequest) (*types.QueryAllCctxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var sends []*types.CrossChainTx
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sendStore := prefix.NewStore(store, types.KeyPrefix(types.SendKey))

	pageRes, err := query.Paginate(sendStore, req.Pagination, func(key []byte, value []byte) error {
		var send types.CrossChainTx
		if err := k.cdc.Unmarshal(value, &send); err != nil {
			return err
		}
		sends = append(sends, &send)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCctxResponse{CrossChainTx: sends, Pagination: pageRes}, nil
}

func (k Keeper) Cctx(c context.Context, req *types.QueryGetCctxRequest) (*types.QueryGetCctxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCrossChainTx(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCctxResponse{CrossChainTx: &val}, nil
}

func (k Keeper) CctxAllPending(c context.Context, req *types.QueryAllCctxPendingRequest) (*types.QueryAllCctxPendingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	//sends := k.GetAllCctxByStatuses(ctx, []types.CctxStatus{types.CctxStatus_PendingOutbound, types.CctxStatus_PendingRevert})
	sends := make([]*types.CrossChainTx, 0)
	pendingTxQueue, found := k.GetPendingTxQueue(ctx, req.ChainId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	startIndex := pendingTxQueue.Head
	// FIXME: use pagination or limit
	for i := 0; i < 100 && i < int(pendingTxQueue.Count); i++ {
		cctx, found := k.GetCrossChainTx(ctx, startIndex)
		if !found {
			return nil, status.Error(codes.Internal, "cctxIndex not found")
		}
		sends = append(sends, &cctx)
		pendingTx, found := k.GetPendingTx(ctx, cctx.Index)
		if !found {
			return nil, status.Error(codes.Internal, "pendingTx not found")
		}
		startIndex = pendingTx.Next
	}

	return &types.QueryAllCctxPendingResponse{CrossChainTx: sends}, nil
}

func (k Keeper) CreateNewCCTX(ctx sdk.Context, msg *types.MsgVoteOnObservedInboundTx, index string, s types.CctxStatus, senderChain, receiverChain *common.Chain) types.CrossChainTx {
	if msg.TxOrigin == "" {
		msg.TxOrigin = msg.Sender
	}
	inboundParams := &types.InboundTxParams{
		Sender:                          msg.Sender,
		SenderChainId:                   senderChain.ChainId,
		TxOrigin:                        msg.TxOrigin,
		Asset:                           msg.Asset,
		Amount:                          msg.Amount,
		CoinType:                        msg.CoinType,
		InboundTxObservedHash:           msg.InTxHash,
		InboundTxObservedExternalHeight: msg.InBlockHeight,
		InboundTxFinalizedZetaHeight:    0,
		InboundTxBallotIndex:            index,
	}

	outBoundParams := &types.OutboundTxParams{
		Receiver:                         msg.Receiver,
		ReceiverChainId:                  receiverChain.ChainId,
		OutboundTxHash:                   "",
		OutboundTxTssNonce:               0,
		OutboundTxGasLimit:               msg.GasLimit,
		OutboundTxGasPrice:               "",
		OutboundTxBallotIndex:            "",
		OutboundTxObservedExternalHeight: 0,
		CoinType:                         msg.CoinType, // FIXME: is this correct?
		Amount:                           sdk.NewUint(0),
	}
	status := &types.Status{
		Status:              s,
		StatusMessage:       "",
		LastUpdateTimestamp: ctx.BlockHeader().Time.Unix(),
	}
	newCctx := types.CrossChainTx{
		Creator:          msg.Creator,
		Index:            index,
		ZetaFees:         math.ZeroUint(),
		RelayedMessage:   msg.Message,
		CctxStatus:       status,
		InboundTxParams:  inboundParams,
		OutboundTxParams: []*types.OutboundTxParams{outBoundParams},
	}
	return newCctx
}
