/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package account

import (
	"cftelescope/cmd/account/intel"

	"github.com/spf13/cobra"
)

var IntelCmd = &cobra.Command{
	Use:   "intel",
	Short: "intel commands for account",
	Long:  `intel commands for the specified account`,
}

func init() {

	IntelCmd.AddCommand(intel.IntelASNGetCmd)
	IntelCmd.AddCommand(intel.IntelDomainHistoryGetCmd)
	IntelCmd.AddCommand(intel.IntelDNSListCmd)
	IntelCmd.AddCommand(intel.IntelWhoisGetCmd)
}
