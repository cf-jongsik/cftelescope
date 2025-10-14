/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package snippets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/snippets"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SnippetsGetCmd represents the snippets command
var SnippetsGetCmd = &cobra.Command{
	Use:   "get <script_name>",
	Short: "Get snippets",
	Long:  `Get a snippet in the account`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		scriptName := args[0]
		res, err := cf.Snippets.Get(ctx, scriptName, snippets.SnippetGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}

		binary, err := json.Marshal(res)
		if err != nil {
			log.Println("marshal dynamic result:", err)
			return
		}
		fmt.Println(string(binary))
	},
}
