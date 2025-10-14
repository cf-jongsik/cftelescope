package rums

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rum"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RumSiteInfoGetCmd represents the site-info command
var RumSiteInfoGetCmd = &cobra.Command{
	Use:   "get <site_tag>",
	Short: "Get RUM site",
	Long:  "Get a specific RUM site in the account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		siteTag := args[0]
		response, err := cf.RUM.SiteInfo.Get(ctx, siteTag, rum.SiteInfoGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
