package logpush

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/logpush"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// LogpushJobsGetCmd represents the jobs command
var LogpushJobsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get logpush jobs",
	Long:  "get logpush jobs for specified job id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		jobIdStr := args[0]
		jobId, err := strconv.ParseInt(jobIdStr, 10, 64)
		if err != nil {
			log.Fatalln(err)
			return
		}
		response, err := cf.Logpush.Jobs.Get(ctx, jobId, logpush.JobGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
