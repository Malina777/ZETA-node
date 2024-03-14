package keeper

import (
	"context"
	"fmt"

	cosmoserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observerkeeper "github.com/zeta-chain/zetacore/x/observer/keeper"
)

// VoteOnObservedOutboundTx casts a vote on an outbound transaction observed on a connected chain (after
// it has been broadcasted to and finalized on a connected chain). If this is
// the first vote, a new ballot is created. When a threshold of votes is
// reached, the ballot is finalized. When a ballot is finalized, the outbound
// transaction is processed.
//
// If the observation is successful, the difference between zeta burned
// and minted is minted by the bank module and deposited into the module
// account.
//
// If the observation is unsuccessful, the logic depends on the previous
// status.
//
// If the previous status was `PendingOutbound`, a new revert transaction is
// created. To cover the revert transaction fee, the required amount of tokens
// submitted with the CCTX are swapped using a Uniswap V2 contract instance on
// ZetaChain for the ZRC20 of the gas token of the receiver chain. The ZRC20
// tokens are then
// burned. The nonce is updated. If everything is successful, the CCTX status is
// changed to `PendingRevert`.
//
// If the previous status was `PendingRevert`, the CCTX is aborted.
//
// ```mermaid
// stateDiagram-v2
//
//	state observation <<choice>>
//	state success_old_status <<choice>>
//	state fail_old_status <<choice>>
//	PendingOutbound --> observation: Finalize outbound
//	observation --> success_old_status: Observation succeeded
//	success_old_status --> Reverted: Old status is PendingRevert
//	success_old_status --> OutboundMined: Old status is PendingOutbound
//	observation --> fail_old_status: Observation failed
//	fail_old_status --> PendingRevert: Old status is PendingOutbound
//	fail_old_status --> Aborted: Old status is PendingRevert
//	PendingOutbound --> Aborted: Finalize outbound error
//
// ```
//
// Only observer validators are authorized to broadcast this message.
func (k msgServer) VoteOnObservedOutboundTx(goCtx context.Context, msg *types.MsgVoteOnObservedOutboundTx) (*types.MsgVoteOnObservedOutboundTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if CCTX exists and if the nonce matches
	cctx, found := k.GetCrossChainTx(ctx, msg.CctxHash)
	if !found {
		return nil, cosmoserrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CCTX %s does not exist", msg.CctxHash))
	}
	if cctx.GetCurrentOutTxParam().OutboundTxTssNonce != msg.OutTxTssNonce {
		return nil, cosmoserrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("OutTxTssNonce %d does not match CCTX OutTxTssNonce %d", msg.OutTxTssNonce, cctx.GetCurrentOutTxParam().OutboundTxTssNonce))
	}

	// get ballot index
	ballotIndex := msg.Digest()

	// vote on outbound ballot
	isFinalizingVote, isNew, ballot, observationChain, err := k.zetaObserverKeeper.VoteOnOutboundBallot(
		ctx,
		ballotIndex,
		msg.OutTxChain,
		msg.Status,
		msg.Creator)
	if err != nil {
		return nil, err
	}
	// if the ballot is new, set the index to the CCTX
	if isNew {
		observerkeeper.EmitEventBallotCreated(ctx, ballot, msg.ObservedOutTxHash, observationChain)
		// Set this the first time when the ballot is created
		// The ballot might change if there are more votes in a different outbound ballot for this cctx hash
		cctx.GetCurrentOutTxParam().OutboundTxBallotIndex = ballotIndex
	}

	// if not finalized commit state here
	if !isFinalizingVote {
		return &types.MsgVoteOnObservedOutboundTxResponse{}, nil
	}

	_, found = k.zetaObserverKeeper.GetTSS(ctx)
	if !found {
		return nil, types.ErrCannotFindTSSKeys
	}
	// if ballot successful, the value received should be the out tx amount
	err = SetOutboundValues(ctx, &cctx, *msg, ballot.BallotStatus)
	if err != nil {
		return nil, err
	}
	// Fund the gas stability pool with the remaining funds
	k.FundStabilityPool(ctx, &cctx)

	err = k.ProcessOutbound(ctx, &cctx, ballot.BallotStatus, msg.ValueReceived.String())
	if err != nil {
		k.SaveFailedOutBound(ctx, &cctx, msg.ValueReceived.String())
		return &types.MsgVoteOnObservedOutboundTxResponse{}, nil
	}
	k.SaveSuccessfulOutBound(ctx, &cctx, msg.ValueReceived.String())
	return &types.MsgVoteOnObservedOutboundTxResponse{}, nil
}
