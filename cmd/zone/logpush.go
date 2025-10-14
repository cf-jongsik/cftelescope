package zone

import (
	logpushPkg "cftelescope/cmd/zone/logpush"

	"github.com/spf13/cobra"
)

var LogpushCmd = &cobra.Command{
	Use:   "logpush",
	Short: "Logpush",
	Long:  "zone",
}

func init() {
	LogpushCmd.AddCommand(logpushPkg.LogpushEdgeGetCmd)
	LogpushCmd.AddCommand(logpushPkg.LogpushJobsListCmd)
	LogpushCmd.AddCommand(logpushPkg.LogpushJobsGetCmd)
}
