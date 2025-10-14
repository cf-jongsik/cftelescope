package logs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/logs"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LogsRayIDGetCmd represents the rayid command
var LogsRayIDGetCmd = &cobra.Command{
	Use:   "rayid",
	Short: "Logs.RayID.Get for Logs.RayID.Get",
	Long:  "zone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		rayID := args[0]
		response, err := cf.Logs.RayID.Get(ctx, rayID, logs.RayIDGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response)
	},
}
