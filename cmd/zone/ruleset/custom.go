/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package ruleset

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rulesets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// custom ruleset represents the rulesets command
var CustomCmd = &cobra.Command{
	Use:   "custom rulesets [cloudflare_zone_id]",
	Short: "custom rulesets commands for a zone",
	Long:  `custom rulesets commands for the specified zone`,
}

var CustomListCmd = &cobra.Command{
	Use:   "list [cloudflare_zone_id] [ruleset-id]",
	Short: "List custom rulesets for a zone",
	Long:  `List all custom rulesets for the specified zone`,
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.Rulesets.ListAutoPaging(ctx, rulesets.RulesetListParams{ZoneID: cloudflare.F(zoneID)})
		for res.Next() {
			if res.Current().Kind != "managed" {
				fmt.Println(res.Current().JSON.RawJSON())
			}
		}
	},
}

func init() {
	CustomCmd.AddCommand(CustomListCmd)
}
