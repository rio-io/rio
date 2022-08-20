package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"rio/x/rio/types"
)

var _ = strconv.Itoa(0)

func CmdSendCert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-cert [to] [cert-type] [title] [description]",
		Short: "Broadcast message send-cert",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTo := args[0]
			argCertType := args[1]
			argTitle := args[2]
			argDescription := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendCert(
				clientCtx.GetFromAddress().String(),
				argTo,
				argCertType,
				argTitle,
				argDescription,
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
