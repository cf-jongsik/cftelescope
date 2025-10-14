package account

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AccountsAccountProfileGetCmd represents the account-profile command
var AccountsAccountProfileGetCmd = &cobra.Command{
	Use:   "profile",
	Short: "get account profile",
	Long:  "get account profile for the specified account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Accounts.AccountProfile.Get(ctx, accounts.AccountProfileGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
