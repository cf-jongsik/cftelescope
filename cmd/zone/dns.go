/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package zone

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DnsListCmd represents the dns command
var DnsListCmd = &cobra.Command{
	Use:   "dns [cloudflare_zone_id]",
	Short: "List DNS records for a zone",
	Long:  `List all DNS records for the specified zone`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		zoneID := viper.GetString("cloudflare_zone_id")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.DNS.Records.ListAutoPaging(ctx, dns.RecordListParams{ZoneID: cloudflare.F(zoneID)})
		for res.Next() {
			fmt.Println(res.Current().JSON.RawJSON())
		}
	},
}
