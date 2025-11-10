/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package aruleset

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

// managed ruleset represents the rulesets command
var ManagedCmd = &cobra.Command{
	Use:   "managed rulesets [cloudflare_zone_id]",
	Short: "managed rulesets commands for a zone",
	Long:  `managed rulesets commands for the specified zone`,
}

var ManagedListCmd = &cobra.Command{
	Use:   "list [cloudflare_zone_id] [ruleset-id]",
	Short: "List managed rulesets for a zone",
	Long:  `List all managed rulesets for the specified zone`,
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.Rulesets.ListAutoPaging(ctx, rulesets.RulesetListParams{AccountID: cloudflare.F(accountID)})
		for res.Next() {
			if res.Current().Kind == "managed" {
				fmt.Println(res.Current().JSON.RawJSON())
			}
		}
	},
}

func init() {
	ManagedCmd.AddCommand(ManagedListCmd)
}
