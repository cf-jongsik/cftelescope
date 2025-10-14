package stream

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/stream"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// StreamGetCmd represents the site-info command
var StreamGetCmd = &cobra.Command{
	Use:   "get <stream_tag>",
	Short: "Get Stream",
	Long:  "Get a specific Stream in the account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		streamTag := args[0]
		response, err := cf.Stream.Get(ctx, streamTag, stream.StreamGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
