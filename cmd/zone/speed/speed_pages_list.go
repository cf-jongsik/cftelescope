package speed

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/speed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SpeedPagesListCmd represents the pages list command
var SpeedPagesListCmd = &cobra.Command{
	Use:   "results",
	Short: "List speed test results",
	Long:  "List speed test results for the specified zone",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response := cf.Speed.Pages.ListAutoPaging(ctx, speed.PageListParams{ZoneID: cloudflare.F(zoneID)})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}
	},
}
