/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package zone

import (
	cachePkg "cftelescope/cmd/zone/cache"

	"github.com/spf13/cobra"
)

// CacheCmd represents the cache command
var CacheCmd = &cobra.Command{
	Use:   "cache [cloudflare_zone_id]",
	Short: "Cache commands for a zone",
	Long:  `Cache commands for the specified zone`,
}

func init() {
	CacheCmd.AddCommand(cachePkg.CacheSmartTieredCacheGetCmd)
	CacheCmd.AddCommand(cachePkg.CacheRegionalTieredCacheGetCmd)
	CacheCmd.AddCommand(cachePkg.CacheCacheReserveGetCmd)
}
