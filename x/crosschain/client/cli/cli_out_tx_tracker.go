package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"strconv"
)

func CmdListOutTxTracker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-out-tx-tracker",
		Short: "list all OutTxTracker",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOutTxTrackerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OutTxTrackerAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOutTxTracker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-out-tx-tracker [index]",
		Short: "shows a OutTxTracker",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetOutTxTrackerRequest{
				Index: argIndex,
			}

			res, err := queryClient.OutTxTracker(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// Transaction CLI /////////////////////////

func CmdAddToWatchList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-to-out-tx-tracker [chain] [nonce] [tx-hash]",
		Short: "Add a out-tx-tracker",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChain := args[0]
			argNonce, _ := strconv.ParseInt(args[1], 10, 64)
			argTxHash := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddToOutTxTracker(
				clientCtx.GetFromAddress().String(),
				argChain,
				uint64(argNonce),
				argTxHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveFromWatchList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-from-out-tx-tracker [chain] [nonce]",
		Short: "Remove a out-tx-tracker",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChain := args[0]
			argNonce, _ := strconv.ParseInt(args[1], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveFromOutTxTracker(
				clientCtx.GetFromAddress().String(),
				argChain,
				uint64(argNonce),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
