/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package snippets

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/snippets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SnippetsListCmd represents the snippets command
var SnippetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List snippets",
	Long:  `List all snippets in the account`,
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.Snippets.ListAutoPaging(ctx, snippets.SnippetListParams{ZoneID: cloudflare.F(zoneID)})
		for res.Next() {
			fmt.Println(res.Current().JSON.RawJSON())
		}
	},
}
