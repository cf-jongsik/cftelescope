package account

import (
	"cftelescope/cmd/account/stream"

	"github.com/spf13/cobra"
)

// StreamCmd represents the stream command
var StreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "stream commends",
	Long:  "stream commends for specified account",
}

func init() {
	StreamCmd.AddCommand(stream.StreamGetCmd)
	StreamCmd.AddCommand(stream.StreamListCmd)
}
