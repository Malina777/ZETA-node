package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zeta-chain/zetacore/common"
)

const TypeMsgDeployFungibleCoinZRC20 = "deploy_fungible_coin_zrc_20"

var _ sdk.Msg = &MsgDeployFungibleCoinZRC20{}

func NewMsgDeployFungibleCoinZRC20(creator string, ERC20 string, foreignChain string, decimals uint32, name string, symbol string, coinType common.CoinType) *MsgDeployFungibleCoinZRC20 {
	return &MsgDeployFungibleCoinZRC20{
		Creator:      creator,
		ERC20:        ERC20,
		ForeignChain: foreignChain,
		Decimals:     decimals,
		Name:         name,
		Symbol:       symbol,
		CoinType:     coinType,
	}
}

func (msg *MsgDeployFungibleCoinZRC20) Route() string {
	return RouterKey
}

func (msg *MsgDeployFungibleCoinZRC20) Type() string {
	return TypeMsgDeployFungibleCoinZRC20
}

func (msg *MsgDeployFungibleCoinZRC20) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeployFungibleCoinZRC20) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeployFungibleCoinZRC20) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
