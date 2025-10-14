/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package zone

import (
	"cftelescope/cmd/zone/ruleset"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rulesets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RulesetsCmd = &cobra.Command{
	Use:   "rulesets [cloudflare_zone_id]",
	Short: "rulesets commands for a zone",
	Long:  `rulesets commands for the specified zone`,
}

// RulesetsListCmd represents the rulesets command
var RulesetsListCmd = &cobra.Command{
	Use:   "list [cloudflare_zone_id]",
	Short: "List rulesets for a zone",
	Long:  `List all rulesets for the specified zone`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		zoneID := viper.GetString("cloudflare_zone_id")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.Rulesets.ListAutoPaging(ctx, rulesets.RulesetListParams{ZoneID: cloudflare.F(zoneID)})
		for res.Next() {
			fmt.Println(res.Current().JSON.RawJSON())
		}
	},
}

var RulesetsGetCmd = &cobra.Command{
	Use:   "get [cloudflare_zone_id] [ruleset-id]",
	Short: "Get ruleset for a zone",
	Long:  `Get a specific ruleset for the specified zone`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		zoneID := viper.GetString("cloudflare_zone_id")
		rulesetID := viper.GetString("ruleset-id")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res, err := cf.Rulesets.Get(ctx, rulesetID, rulesets.RulesetGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(res.JSON.RawJSON())
	},
}

func init() {
	RulesetsCmd.PersistentFlags().StringP("ruleset-id", "r", "", "ruleset id")
	viper.BindPFlag("ruleset-id", RulesetsCmd.PersistentFlags().Lookup("ruleset-id"))

	RulesetsCmd.AddCommand(RulesetsListCmd)
	RulesetsCmd.AddCommand(RulesetsGetCmd)
	RulesetsCmd.AddCommand(ruleset.ManagedCmd)
	RulesetsCmd.AddCommand(ruleset.CustomCmd)
}
