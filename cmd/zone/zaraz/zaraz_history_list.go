package zaraz

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zaraz"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ZarazHistoryListCmd represents the history command
var ZarazHistoryListCmd = &cobra.Command{
	Use:   "history",
	Short: "Zaraz.History.List for Zaraz.History.List",
	Long:  "zone",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response := cf.Zaraz.History.ListAutoPaging(ctx, zaraz.HistoryListParams{ZoneID: cloudflare.F(zoneID)})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}
	},
}
