package cli

import (
	"strconv"

	"github.com/alice/checkers/x/oddnum/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [from] [to] [amount]",
		Short: "Broadcast message transfer",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFrom := args[0]
			argTo := args[1]
			argAmount, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransfer(
				clientCtx.GetFromAddress().String(),
				argFrom,
				argTo,
				argAmount,
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
