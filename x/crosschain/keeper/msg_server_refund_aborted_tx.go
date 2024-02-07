package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"golang.org/x/net/context"
)

func (k msgServer) RefundAbortedCCTX(goCtx context.Context, msg *types.MsgRefundAbortedCCTX) (*types.MsgRefundAbortedCCTXResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if authorized
	if msg.Creator != k.zetaObserverKeeper.GetParams(ctx).GetAdminPolicyAccount(observertypes.Policy_Type_group2) {
		return nil, observertypes.ErrNotAuthorized
	}
	// check if the cctx exists
	cctx, found := k.GetCrossChainTx(ctx, msg.CctxIndex)
	if !found {
		return nil, types.ErrCannotFindCctx
	}
	// check if the cctx is aborted
	if cctx.CctxStatus.Status != types.CctxStatus_Aborted {
		return nil, errorsmod.Wrap(types.ErrInvalidStatus, "CCTX is not aborted")
	}
	// check if the cctx is not refunded
	if cctx.IsRefunded {
		return nil, errorsmod.Wrap(types.ErrUnableProcessRefund, "CCTX is already refunded")
	}

	// refund the amount
	if common.IsEVMChain(cctx.InboundTxParams.SenderChainId) {
		err := k.RefundAbortedAmountOnZetaChainForEvmChain(ctx, cctx)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrUnableProcessRefund, err.Error())
		}
	} else if common.IsBitcoinChain(cctx.InboundTxParams.SenderChainId) {
		err := k.RefundAbortedAmountOnZetaChainForBitcoinChain(ctx, cctx, msg.ReceiverBtcRefund)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrUnableProcessRefund, err.Error())
		}
	}

	cctx.IsRefunded = true
	k.SetCrossChainTx(ctx, cctx)
	if cctx.GetCurrentOutTxParam().CoinType == common.CoinType_Zeta {
		k.RemoveZetaAbortedAmount(ctx, cctx.GetCurrentOutTxParam().Amount)
	}
	return &types.MsgRefundAbortedCCTXResponse{}, nil
}