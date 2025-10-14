package rums

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rum"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RumSiteInfoListCmd represents the site-info command
var RumSiteInfoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List RUM sites",
	Long:  "List all RUM sites in the account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response := cf.RUM.SiteInfo.ListAutoPaging(ctx, rum.SiteInfoListParams{AccountID: cloudflare.F(accountID)})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}
	},
}
