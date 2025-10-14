/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package pages

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PagesListCmd represents the pages command
var PagesProjectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Pages projects",
	Long:  `List all Pages projects in the account`,
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		res := cf.Pages.Projects.ListAutoPaging(ctx, pages.ProjectListParams{AccountID: cloudflare.F(accountID)})
		for res.Next() {
			fmt.Println(res.Current().JSON.RawJSON())
		}
	},
}
