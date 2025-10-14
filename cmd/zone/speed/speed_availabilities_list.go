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

// SpeedAvailabilitiesListCmd represents the availabilities list command
var SpeedAvailabilitiesListCmd = &cobra.Command{
	Use:   "quotas",
	Short: "List speed test quotas",
	Long:  "List speed test quotas for the specified zone",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Speed.Availabilities.List(ctx, speed.AvailabilityListParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
