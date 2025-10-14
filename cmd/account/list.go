/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package account

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rules"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListListCmd represents the list command
var ListListCmd = &cobra.Command{
	Use:   "lists",
	Short: "List list items for account",
	Long:  `List all list items for the specified account`,
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		list := cf.Rules.Lists.ListAutoPaging(ctx, rules.ListListParams{AccountID: cloudflare.F(accountID)})
		for list.Next() {
			fmt.Println(list.Current().JSON.RawJSON())
		}
	},
}
