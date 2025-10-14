package zaraz

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zaraz"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ZarazWorkflowGetCmd represents the workflow command
var ZarazWorkflowGetCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Zaraz.Workflow.Get for Zaraz.Workflow.Get",
	Long:  "zone",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Zaraz.Workflow.Get(ctx, zaraz.WorkflowGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response)
	},
}
