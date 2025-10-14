package zone

import (
	logsPkg "cftelescope/cmd/zone/logs"

	"github.com/spf13/cobra"
)

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "logs",
	Long:  "logs",
}

func init() {
	LogsCmd.AddCommand(logsPkg.LogsControlRetentionGetCmd)
	LogsCmd.AddCommand(logsPkg.LogsRayIDGetCmd)
	LogsCmd.AddCommand(logsPkg.LogsReceivedGetCmd)
}
