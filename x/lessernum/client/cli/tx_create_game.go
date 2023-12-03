package cli

import (
	"github.com/spf13/cast"
	"strconv"

	"github.com/alice/checkers/x/lessernum/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [player-1] [player-2] [wager]",
		Short: "Broadcast message createGame",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPlayer1 := args[0]
			argPlayer2 := args[1]
			argWager, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGame(
				clientCtx.GetFromAddress().String(),
				argPlayer1,
				argPlayer2,
				argWager,
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
