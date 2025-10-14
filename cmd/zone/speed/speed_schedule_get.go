package speed

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/speed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SpeedScheduleGetCmd represents the schedule command
var SpeedScheduleGetCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Get speed test schedule",
	Long:  "Get speed test schedule for the specified zone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		scheduleID := args[0]

		response, err := cf.Speed.Schedule.Get(ctx, scheduleID, speed.ScheduleGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
