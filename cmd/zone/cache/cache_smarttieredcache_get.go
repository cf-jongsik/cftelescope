package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/cache"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CacheSmartTieredCacheGetCmd represents the smart-tiered-cache command
var CacheSmartTieredCacheGetCmd = &cobra.Command{
	Use:   "smart",
	Short: "Cache Smart Tiered Cache",
	Long:  "Cache Smart Tiered Cache",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Cache.SmartTieredCache.Get(ctx, cache.SmartTieredCacheGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
