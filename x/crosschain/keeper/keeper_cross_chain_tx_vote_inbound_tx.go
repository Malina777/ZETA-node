package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	zetaObserverTypes "github.com/zeta-chain/zetacore/x/observer/types"
)

// FIXME: use more specific error types & codes
func (k msgServer) VoteOnObservedInboundTx(goCtx context.Context, msg *types.MsgVoteOnObservedInboundTx) (*types.MsgVoteOnObservedInboundTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	observationType := zetaObserverTypes.ObservationType_InBoundTx
	if !k.IsInboundAllowed(ctx) {
		return nil, types.ErrNotEnoughPermissions
	}
	// GetChainFromChainID makes sure we are getting only supported chains , if a chain support has been turned on using gov proposal, this function returns nil
	observationChain := k.zetaObserverKeeper.GetParams(ctx).GetChainFromChainID(msg.SenderChainId)
	if observationChain == nil {
		return nil, sdkerrors.Wrap(types.ErrUnsupportedChain, fmt.Sprintf("ChainID %d, Observation %s", msg.SenderChainId, observationType.String()))
	}
	receiverChain := k.zetaObserverKeeper.GetParams(ctx).GetChainFromChainID(msg.ReceiverChain)
	if receiverChain == nil {
		return nil, sdkerrors.Wrap(types.ErrUnsupportedChain, fmt.Sprintf("ChainID %d, Observation %s", msg.ReceiverChain, observationType.String()))
	}
	// IsAuthorized does various checks against the list of observer mappers
	ok, err := k.IsAuthorized(ctx, msg.Creator, observationChain)
	if !ok {
		return nil, err
	}

	index := msg.Digest()
	// Add votes and Set Ballot
	// GetBallot checks against the supported chains list before querying for Ballot
	ballot, isNew, err := k.GetBallot(ctx, index, observationChain, observationType)
	if err != nil {
		return nil, err
	}
	if isNew {
		EmitEventBallotCreated(ctx, ballot, msg.InTxHash, observationChain.String())
	}
	// AddVoteToBallot adds a vote and sets the ballot
	ballot, err = k.AddVoteToBallot(ctx, ballot, msg.Creator, zetaObserverTypes.VoteType_SuccessObservation)
	if err != nil {
		return nil, err
	}
	// CheckIfBallotIsFinalized checks status and sets the ballot if finalized

	ballot, isFinalized := k.CheckIfBallotIsFinalized(ctx, ballot)
	if !isFinalized {
		return &types.MsgVoteOnObservedInboundTxResponse{}, nil
	}

	// ******************************************************************************
	// below only happens when ballot is finalized: exactly when threshold vote is in
	// ******************************************************************************

	// Inbound Ballot has been finalized , Create CCTX
	cctx := k.CreateNewCCTX(ctx, msg, index, types.CctxStatus_PendingInbound, observationChain, receiverChain)
	// FinalizeInbound updates CCTX Prices and Nonce
	// Aborts is any of the updates fail

	if receiverChain.IsZetaChain() {
		cachedCtx, write := ctx.CacheContext()
		err = k.HandleEVMDeposit(cachedCtx, &cctx, *msg, observationChain)
		if err != nil {
			EmitEventInboundFailed(ctx, &cctx)
			cctx.CctxStatus.ChangeStatus(&ctx, types.CctxStatus_Aborted, err.Error(), cctx.LogIdentifierForCCTX())
			k.SetCrossChainTx(ctx, cctx)
			return &types.MsgVoteOnObservedInboundTxResponse{}, nil
		}
		write()
		cctx.CctxStatus.ChangeStatus(&ctx, types.CctxStatus_OutboundMined, "First half of EVM transfer Completed", cctx.LogIdentifierForCCTX())
	} else { // Cross Chain SWAP
		cachedCtx, write := ctx.CacheContext()
		err = k.FinalizeInbound(cachedCtx, &cctx, *receiverChain, len(ballot.VoterList))
		if err != nil {
			EmitEventInboundFailed(ctx, &cctx)
			cctx.CctxStatus.ChangeStatus(&ctx, types.CctxStatus_Aborted, err.Error(), cctx.LogIdentifierForCCTX())
			k.SetCrossChainTx(ctx, cctx)
			return &types.MsgVoteOnObservedInboundTxResponse{}, nil
		}
		write()
		cctx.CctxStatus.ChangeStatus(&ctx, types.CctxStatus_PendingOutbound, "Status Changed to Pending Outbound", cctx.LogIdentifierForCCTX())
	}
	EmitEventInboundFinalized(ctx, &cctx)
	EmitEventBallotFinalized(ctx, ballot, observationChain.String())
	k.SetCrossChainTx(ctx, cctx)
	return &types.MsgVoteOnObservedInboundTxResponse{}, nil
}

// TODO: is LastBlockHeight needed?
func (k msgServer) FinalizeInbound(ctx sdk.Context, cctx *types.CrossChainTx, receiveChain common.Chain, numberofobservers int) error {
	cctx.InboundTxParams.InboundTxFinalizedZetaHeight = uint64(ctx.BlockHeader().Height)
	//k.UpdateLastBlockHeight(ctx, cctx)

	err := k.UpdatePrices(ctx, receiveChain.ChainId, cctx)
	if err != nil {
		return err
	}
	err = k.UpdateNonce(ctx, receiveChain.ChainId, cctx)
	if err != nil {
		return err
	}
	return nil
}

//func (k msgServer) UpdateLastBlockHeight(ctx sdk.Context, msg *types.CrossChainTx) {
//	lastblock, isFound := k.GetLastBlockHeight(ctx, msg.InboundTxParams.SenderChain)
//	if !isFound {
//		lastblock = types.LastBlockHeight{
//			Creator:           msg.Creator,
//			Index:             msg.InboundTxParams.SenderChain, // ?
//			Chain:             msg.InboundTxParams.SenderChain,
//			LastSendHeight:    msg.InboundTxParams.InboundTxObservedExternalHeight,
//			LastReceiveHeight: 0,
//		}
//	} else {
//		lastblock.LastSendHeight = msg.InboundTxParams.InboundTxObservedExternalHeight
//	}
//	k.SetLastBlockHeight(ctx, lastblock)
//}
