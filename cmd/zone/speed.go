package zone

import (
	speedPkg "cftelescope/cmd/zone/speed"

	"github.com/spf13/cobra"
)

// SpeedCmd represents the speed command
var SpeedCmd = &cobra.Command{
	Use:   "speed",
	Short: "speed commends",
	Long:  "speed commends for specified zone",
}

func init() {
	SpeedCmd.AddCommand(speedPkg.SpeedPagesListCmd)
	SpeedCmd.AddCommand(speedPkg.SpeedAvailabilitiesListCmd)
	SpeedCmd.AddCommand(speedPkg.SpeedScheduleGetCmd)
}
