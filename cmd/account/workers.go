/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package account

import (
	"cftelescope/cmd/account/workers"

	"github.com/spf13/cobra"
)

// WorkersCmd represents the workers command
var WorkersCmd = &cobra.Command{
	Use:   "workers",
	Short: "workers commands",
	Long:  `workers commands`,
}

func init() {
	WorkersCmd.AddCommand(workers.WorkersListCmd)
	WorkersCmd.AddCommand(workers.WorkersGetCmd)
}
