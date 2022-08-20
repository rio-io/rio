package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"rio/x/rio/types"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdCreateResume() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-resume [certs] [avatar-url] [name] [description]",
		Short: "Broadcast message create-resume",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCastCerts := strings.Split(args[0], listSeparator)
			argCerts := make([]uint64, len(argCastCerts))
			for i, arg := range argCastCerts {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argCerts[i] = value
			}
			argAvatarUrl := args[1]
			argName := args[2]
			argDescription := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateResume(
				clientCtx.GetFromAddress().String(),
				argCerts,
				argAvatarUrl,
				argName,
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
