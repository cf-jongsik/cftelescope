/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// WorkersListCmd represents the workers command
var WorkersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workers",
	Long:  `List all workers in the account`,
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.Workers.Scripts.ListAutoPaging(ctx, workers.ScriptListParams{AccountID: cloudflare.F(accountID)})
		for res.Next() {
			fmt.Println(res.Current().JSON.RawJSON())
		}
	},
}
