package account

import (
	"cftelescope/cmd/account/rums"

	"github.com/spf13/cobra"
)

var RUMsCmd = &cobra.Command{
	Use:   "rums",
	Short: "RUM commands for account",
	Long:  `RUM commands for the specified account`,
}

func init() {
	RUMsCmd.AddCommand(rums.RumSiteInfoListCmd)
	RUMsCmd.AddCommand(rums.RumSiteInfoGetCmd)
}
