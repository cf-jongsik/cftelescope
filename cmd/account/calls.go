package account

import (
	"cftelescope/cmd/account/calls"

	"github.com/spf13/cobra"
)

var CallsCmd = &cobra.Command{
	Use:   "calls",
	Short: "call commends",
	Long:  "call commends",
}

var CallsTurnCmd = &cobra.Command{
	Use:   "turn",
	Short: "turn commends",
	Long:  "turn commends",
}

var CallsSFUCmd = &cobra.Command{
	Use:   "sfu",
	Short: "sfu commends",
	Long:  "sfu commends",
}

func init() {
	CallsCmd.AddCommand(CallsTurnCmd)
	CallsCmd.AddCommand(CallsSFUCmd)
	CallsSFUCmd.AddCommand(calls.CallsSFUListCmd)
	CallsSFUCmd.AddCommand(calls.CallsSFUGetCmd)
	CallsTurnCmd.AddCommand(calls.CallsTURNListCmd)
	CallsTurnCmd.AddCommand(calls.CallsTURNGetCmd)
}
