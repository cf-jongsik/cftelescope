package zone

import (
	zarazPkg "cftelescope/cmd/zone/zaraz"

	"github.com/spf13/cobra"
)

var ZarazCmd = &cobra.Command{
	Use:   "zaraz",
	Short: "Zaraz",
	Long:  `Zaraz`,
}

func init() {
	ZarazCmd.AddCommand(zarazPkg.ZarazConfigGetCmd)
	ZarazCmd.AddCommand(zarazPkg.ZarazDefaultGetCmd)
	ZarazCmd.AddCommand(zarazPkg.ZarazExportGetCmd)
	ZarazCmd.AddCommand(zarazPkg.ZarazHistoryListCmd)
	ZarazCmd.AddCommand(zarazPkg.ZarazWorkflowGetCmd)
}
