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

// CacheRegionalTieredCacheGetCmd represents the regional-tiered-cache command
var CacheRegionalTieredCacheGetCmd = &cobra.Command{
	Use:   "regional",
	Short: "Cache Regional Tiered Cache",
	Long:  "Cache Regional Tiered Cache",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Cache.RegionalTieredCache.Get(ctx, cache.RegionalTieredCacheGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
