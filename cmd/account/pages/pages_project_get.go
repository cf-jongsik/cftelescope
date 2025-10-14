/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package pages

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var PagesProjectGetCmd = &cobra.Command{
	Use:   "get <project_name>",
	Short: "Get Pages project",
	Args:  cobra.ExactArgs(1),
	Long:  `Get a specific Pages project in the account`,
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		projectName := args[0]
		res, err := cf.Pages.Projects.Get(ctx, projectName, pages.ProjectGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(res.JSON.RawJSON())
	},
}
