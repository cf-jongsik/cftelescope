package d1

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/d1"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// D1DatabaseListCmd represents the database command
var D1DatabaseListCmd = &cobra.Command{
	Use:   "list",
	Short: "list d1 database",
	Long:  "list d1 database for the specified account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response := cf.D1.Database.ListAutoPaging(ctx, d1.DatabaseListParams{AccountID: cloudflare.F(accountID)})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}
	},
}
