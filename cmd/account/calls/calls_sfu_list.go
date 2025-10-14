package calls

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/calls"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CallsSFUListCmd represents the s-f-u command
var CallsSFUListCmd = &cobra.Command{
	Use:   "list",
	Short: "list sfu",
	Long:  "list sfu for the specified account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response := cf.Calls.SFU.ListAutoPaging(ctx, calls.SFUListParams{AccountID: cloudflare.F(accountID)})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}
	},
}
