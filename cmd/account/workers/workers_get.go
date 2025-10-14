/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// WorkersGetCmd represents the workers command
var WorkersGetCmd = &cobra.Command{
	Use:   "get <script_name>",
	Short: "Get workers",
	Long:  `Get all workers in the account`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		scriptName := args[0]
		res, err := cf.Workers.Scripts.Get(ctx, scriptName, workers.ScriptGetParams{AccountID: cloudflare.F(accountID)})
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
