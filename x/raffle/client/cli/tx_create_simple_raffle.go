package cli

import (
	"strconv"

	"raffle/x/raffle/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateSimpleRaffle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-simple-raffle [title] [description] [participant-list-url] [number-of-winners] [number-of-participants]",
		Short: "Broadcast message createSimpleRaffle",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argParticipantListUrl := args[2]
			argNumberOfWinners, err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				return err
			}

			argNumberOfParticipants, err := strconv.ParseUint(args[4], 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSimpleRaffle(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argParticipantListUrl,
				uint32(argNumberOfWinners),
				uint32(argNumberOfParticipants),
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
