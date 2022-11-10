package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"github.com/spf13/cobra"
)

func CmdListFeedist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-feedist",
		Short: "list all feedist",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFeedistRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FeedistAll(context.Background(), params)
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

func CmdShowFeedist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-feedist [index]",
		Short: "shows a feedist",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetFeedistRequest{
				Index: argIndex,
			}

			res, err := queryClient.Feedist(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
