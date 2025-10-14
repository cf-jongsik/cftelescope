package calls

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/calls"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CallsTURNGetCmd represents the t-u-r-n command
var CallsTURNGetCmd = &cobra.Command{
	Use:   "get [keyID]",
	Short: "get turn",
	Long:  "get turn for the specified account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		keyID := args[0]
		response, err := cf.Calls.TURN.Get(ctx, keyID, calls.TURNGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
