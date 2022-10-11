package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
	"github.com/zeta-chain/zetacore/x/zetacore/types"
	zetaObserverTypes "github.com/zeta-chain/zetacore/x/zetaobserver/types"
)

func (k msgServer) VoteOnObservedInboundTx(goCtx context.Context, msg *types.MsgVoteOnObservedInboundTx) (*types.MsgVoteOnObservedInboundTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	observationType := zetaObserverTypes.ObservationType_InBoundTx
	observationChain := zetaObserverTypes.ConvertStringChaintoObservationChain(msg.SenderChain)
	receiverChain := zetaObserverTypes.ConvertStringChaintoObservationChain(msg.ReceiverChain)
	ok, err := k.isAuthorized(ctx, msg.Creator, observationChain, observationType.String())
	if !ok {
		return nil, err
	}
	index := msg.Digest()
	// Add votes and Set Ballot
	var ballot zetaObserverTypes.Ballot
	ballot, found := k.zetaObserverKeeper.GetBallot(ctx, index)
	if !found {
		if !k.zetaObserverKeeper.IsChainSupported(ctx, observationChain) || !k.zetaObserverKeeper.IsChainSupported(ctx, receiverChain) {
			return nil, sdkerrors.Wrap(types.ErrUnsupportedChain, "Receiving chain is not supported")
		}
		observerMapper, _ := k.zetaObserverKeeper.GetObserverMapper(ctx, observationChain, observationType.String())
		threshohold, found := k.zetaObserverKeeper.GetParams(ctx).GetVotingThreshold(observationChain, observationType)
		if !found {
			return nil, errors.Wrap(zetaObserverTypes.ErrSupportedChains, fmt.Sprintf("Thresholds not set for Chain %s and Observation %s", observationChain.String(), observationType))
		}

		ballot = zetaObserverTypes.Ballot{
			Index:            "",
			BallotIdentifier: index,
			VoterList:        zetaObserverTypes.CreateVoterList(observerMapper.ObserverList),
			ObservationType:  observationType,
			BallotThreshold:  threshohold.Threshold,
			BallotStatus:     zetaObserverTypes.BallotStatus_BallotInProgress,
		}
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
	// Inbound Ballot has been finalized , Create CCTX
	// Ballot and CCTX have same index , but different prefix stores
	cctx := k.CreateNewCCTX(ctx, msg, index)
	oldStatus := cctx.CctxStatus.Status
	// Finalize updates CCTX Prices and Nonce , Abort CCTX is any of the updates fail
	err = k.FinalizeInbound(ctx, &cctx, msg.ReceiverChain, len(ballot.VoterList))
	if err != nil {
		cctx.CctxStatus.ChangeStatus(&ctx, types.CctxStatus_Aborted, err.Error(), cctx.LogIdentifierForCCTX())
		ctx.Logger().Error(err.Error())
		k.CctxChangePrefixStore(ctx, cctx, oldStatus)
		return &types.MsgVoteOnObservedInboundTxResponse{}, nil
	}

	cctx.CctxStatus.ChangeStatus(&ctx, types.CctxStatus_PendingOutbound, "Status Changed to Pending Outbound", cctx.LogIdentifierForCCTX())
	k.CctxChangePrefixStore(ctx, cctx, oldStatus)
	return &types.MsgVoteOnObservedInboundTxResponse{}, nil

}

func (k msgServer) FinalizeInbound(ctx sdk.Context, cctx *types.CrossChainTx, receiveChain string, numberofobservers int) error {
	cctx.InBoundTxParams.InBoundTxFinalizedZetaHeight = uint64(ctx.BlockHeader().Height)
	k.UpdateLastBlockHeight(ctx, cctx)
	bftTime := ctx.BlockHeader().Time // we use BFTTime of the current block as random number
	cctx.OutBoundTxParams.Broadcaster = uint64(bftTime.Nanosecond() % numberofobservers)

	err := k.UpdatePrices(ctx, receiveChain, cctx)
	if err != nil {
		return err
	}
	err = k.UpdateNonce(ctx, receiveChain, cctx)
	if err != nil {
		return err
	}
	EmitEventSendFinalized(ctx, cctx)
	return nil
}

func (k msgServer) UpdateLastBlockHeight(ctx sdk.Context, msg *types.CrossChainTx) {
	lastblock, isFound := k.GetLastBlockHeight(ctx, msg.InBoundTxParams.SenderChain)
	if !isFound {
		lastblock = types.LastBlockHeight{
			Creator:           msg.Creator,
			Index:             msg.InBoundTxParams.SenderChain, // ?
			Chain:             msg.InBoundTxParams.SenderChain,
			LastSendHeight:    msg.InBoundTxParams.InBoundTxObservedExternalHeight,
			LastReceiveHeight: 0,
		}
	} else {
		lastblock.LastSendHeight = msg.InBoundTxParams.InBoundTxObservedExternalHeight
	}
	k.SetLastBlockHeight(ctx, lastblock)
}
