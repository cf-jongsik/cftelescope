/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package cmd

import (
	zonePkg "cftelescope/cmd/zone"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// zoneCmd represents the zone command group
var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Cloudflare zones",
	Long:  `Cloudflare zones configs`,
}

var zoneInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get zone details",
	Long:  `Get detailed information about a specific zone`,
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		zone, err := cf.Zones.Get(ctx, zones.ZoneGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		log.Printf("ID: %s\nName: %s\nStatus: %s\nType: %s\n", zone.ID, zone.Name, zone.Status, zone.Type)
	},
}

var zoneListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all zones",
	Long:  `List all zones in the account`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		zones := cf.Zones.ListAutoPaging(ctx, zones.ZoneListParams{})
		for zones.Next() {
			fmt.Println(zones.Current().JSON.RawJSON())
		}
	},
}

func init() {
	// Register the zone command with its subcommands
	rootCmd.AddCommand(zoneCmd)

	zoneCmd.PersistentFlags().StringP("cloudflare_zone_id", "z", "", "Zone ID")
	viper.BindPFlag("cloudflare_zone_id", zoneCmd.PersistentFlags().Lookup("cloudflare_zone_id"))

	// Add top-level subcommands that don't require a zone ID
	zoneCmd.AddCommand(zoneListCmd)
	zoneCmd.AddCommand(zoneInfoCmd)
	zoneCmd.AddCommand(zonePkg.DnsListCmd)
	zoneCmd.AddCommand(zonePkg.SslCmd)
	zoneCmd.AddCommand(zonePkg.CacheCmd)
	zoneCmd.AddCommand(zonePkg.RulesetsCmd)
	zoneCmd.AddCommand(zonePkg.ArgoCmd)
	zoneCmd.AddCommand(zonePkg.SpeedCmd)
	zoneCmd.AddCommand(zonePkg.LogsCmd)
	zoneCmd.AddCommand(zonePkg.LbListCmd)
	zoneCmd.AddCommand(zonePkg.DnsDNSSECGetCmd)
	zoneCmd.AddCommand(zonePkg.LogpushCmd)
	zoneCmd.AddCommand(zonePkg.SnippetsCmd)
	zoneCmd.AddCommand(zonePkg.ZarazCmd)
}
