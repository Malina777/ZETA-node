package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/pkg/chains"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

func TestMsgRemoveChainParams_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgRemoveChainParams
		err  error
	}{
		{
			name: "valid message",
			msg: types.NewMsgRemoveChainParams(
				sample.AccAddress(),
				chains.ExternalChainList()[0].ChainId,
			),
		},
		{
			name: "invalid address",
			msg: types.NewMsgRemoveChainParams(
				"invalid_address",
				chains.ExternalChainList()[0].ChainId,
			),
			err: sdkerrors.ErrInvalidAddress,
		},

		{
			name: "invalid chain ID",
			msg: types.NewMsgRemoveChainParams(
				sample.AccAddress(),
				999,
			),
			err: sdkerrors.ErrInvalidChainID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgRemoveChainParams_GetSigners(t *testing.T) {
	signer := sample.AccAddress()
	tests := []struct {
		name   string
		msg    types.MsgRemoveChainParams
		panics bool
	}{
		{
			name: "valid signer",
			msg: types.MsgRemoveChainParams{
				Creator: signer,
			},
			panics: false,
		},
		{
			name: "invalid signer",
			msg: types.MsgRemoveChainParams{
				Creator: "invalid",
			},
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panics {
				signers := tt.msg.GetSigners()
				require.Equal(t, []sdk.AccAddress{sdk.MustAccAddressFromBech32(signer)}, signers)
			} else {
				require.Panics(t, func() {
					tt.msg.GetSigners()
				})
			}
		})
	}
}

func TestMsgRemoveChainParams_Type(t *testing.T) {
	msg := types.MsgRemoveChainParams{
		Creator: sample.AccAddress(),
	}
	require.Equal(t, types.TypeMsgRemoveChainParams, msg.Type())
}

func TestMsgRemoveChainParams_Route(t *testing.T) {
	msg := types.MsgRemoveChainParams{
		Creator: sample.AccAddress(),
	}
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgRemoveChainParams_GetSignBytes(t *testing.T) {
	msg := types.MsgRemoveChainParams{
		Creator: sample.AccAddress(),
	}
	require.NotPanics(t, func() {
		msg.GetSignBytes()
	})
}
