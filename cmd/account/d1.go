package account

import (
	"cftelescope/cmd/account/d1"

	"github.com/spf13/cobra"
)

var D1Cmd = &cobra.Command{
	Use:   "d1",
	Short: "d1 commands for account",
	Long:  `d1 commands for the specified account`,
}

func init() {
	D1Cmd.AddCommand(d1.D1DatabaseGetCmd)
	D1Cmd.AddCommand(d1.D1DatabaseListCmd)
}
