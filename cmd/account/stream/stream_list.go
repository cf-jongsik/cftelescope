package stream

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/stream"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// StreamListCmd represents the site-info command
var StreamListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Stream",
	Long:  "List all Stream in the account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		list := cf.Stream.ListAutoPaging(ctx, stream.StreamListParams{AccountID: cloudflare.F(accountID)})
		for list.Next() {
			fmt.Println(list.Current().JSON.RawJSON())
		}
	},
}
