package logpush

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/logpush"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LogpushJobsListCmd represents the jobs command
var LogpushJobsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list logpush jobs",
	Long:  "list logpush jobs for specified zone",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response := cf.Logpush.Jobs.ListAutoPaging(ctx, logpush.JobListParams{ZoneID: cloudflare.F(zoneID)})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}

	},
}
