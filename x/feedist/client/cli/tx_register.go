package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterFeedist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [contract] [shares]",
		Short: "Broadcast message register",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContract := args[0]
			argShares := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			shares, err := sdk.NewDecFromStr(argShares)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterFeedist(
				clientCtx.GetFromAddress().String(),
				argContract,
				shares,
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
