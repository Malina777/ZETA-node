//go:build TESTNET
// +build TESTNET

package types

import (
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/zeta-chain/zetacore/common"
)

func GetCoreParams() CoreParamsList {
	params := CoreParamsList{
		CoreParams: []*CoreParams{
			{
				ChainId:                     common.GoerliChain().ChainId,
				ConfirmationCount:           14,
				ZetaTokenContractAddress:    "0xfF8dee1305D6200791e26606a0b04e12C5292aD8",
				ConnectorContractAddress:    "0x851b2446f225266C4EC3cd665f6801D624626c4D",
				Erc20CustodyContractAddress: "0x0e141A7e7C0A7E15E7d22713Fc0a6187515Fa9BF",
				InTxTicker:                  24,
				OutTxTicker:                 3,
				WatchUtxoTicker:             0,
				GasPriceTicker:              5,
			},
			{
				ChainId:                     common.BscTestnetChain().ChainId,
				ConfirmationCount:           15,
				ZetaTokenContractAddress:    "0x33580e10212342d0aA66C9de3F6F6a4AfefA144C",
				ConnectorContractAddress:    "0xcF1B4B432CA02D6418a818044d38b18CDd3682E9",
				Erc20CustodyContractAddress: "0x0e141A7e7C0A7E15E7d22713Fc0a6187515Fa9BF",
				InTxTicker:                  24,
				OutTxTicker:                 3,
				WatchUtxoTicker:             0,
				GasPriceTicker:              5,
			},
			{
				ChainId:                     common.MumbaiChain().ChainId,
				ConfirmationCount:           128,
				ZetaTokenContractAddress:    "0xBaEF590c5Aef9881b0a5C86e18D35432218C64D5",
				ConnectorContractAddress:    "0xED4d7f8cA6252Ccf85A1eFB5444d7dB794ddD328",
				Erc20CustodyContractAddress: "0x0e141A7e7C0A7E15E7d22713Fc0a6187515Fa9BF",
				InTxTicker:                  24,
				OutTxTicker:                 3,
				WatchUtxoTicker:             0,
				GasPriceTicker:              5,
			},
			{
				ChainId:                     common.BaobabChain().ChainId,
				ConfirmationCount:           24,
				ZetaTokenContractAddress:    "0x000080383847bD75F91c168269Aa74004877592f",
				ConnectorContractAddress:    "0x000054d3A0Bc83Ec7808F52fCdC28A96c89F6C5c",
				Erc20CustodyContractAddress: "0x0e141A7e7C0A7E15E7d22713Fc0a6187515Fa9BF",
				InTxTicker:                  24,
				OutTxTicker:                 3,
				WatchUtxoTicker:             0,
				GasPriceTicker:              5,
			},
			{
				ChainId:                     common.ZetaChain().ChainId,
				ConfirmationCount:           3,
				ZetaTokenContractAddress:    "0x2DD9830f8Ac0E421aFF9B7c8f7E9DF6F65DBF6Ea",
				ConnectorContractAddress:    "",
				Erc20CustodyContractAddress: "",
				InTxTicker:                  24,
				OutTxTicker:                 3,
				WatchUtxoTicker:             0,
				GasPriceTicker:              5,
			},
			{
				ChainId:                     common.BtcTestNetChain().ChainId,
				ConfirmationCount:           3,
				ZetaTokenContractAddress:    "0x2DD9830f8Ac0E421aFF9B7c8f7E9DF6F65DBF6Ea",
				ConnectorContractAddress:    "",
				Erc20CustodyContractAddress: "0x0e141A7e7C0A7E15E7d22713Fc0a6187515Fa9BF",
				InTxTicker:                  24,
				OutTxTicker:                 3,
				WatchUtxoTicker:             0,
				GasPriceTicker:              5,
			},
		},
	}
	chainList := common.DefaultChainsList()
	requiredParams := len(chainList)
	availableParams := 0
	for _, chain := range chainList {
		for _, param := range params.CoreParams {
			if chain.ChainId == param.ChainId {
				availableParams++
			}
		}
	}
	if availableParams != requiredParams {
		panic(fmt.Sprintf("Core params are not available for all chains , DefaultChains : %s , CoreParams : %s",
			types.PrettyPrintStruct(chainList), params.String()))
	}
	return params
}
