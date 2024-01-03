package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

func CmdWhiteListERC20() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-erc20 [erc20Address] [chainID] [name] [symbol] [decimals] [gasLimit]",
		Short: "Add a new erc20 address to whitelist",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			erc20Address := args[0]
			chainID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			name := args[2]
			symbol := args[3]
			decimals, err := strconv.ParseUint(args[4], 10, 32)
			if err != nil {
				return err
			}

			gasLimit, err := strconv.ParseInt(args[5], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWhitelistERC20(
				clientCtx.GetFromAddress().String(),
				erc20Address,
				chainID,
				name,
				symbol,
				uint32(decimals),
				gasLimit,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}