package intel

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/intel"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// IntelDomainHistoryGetCmd represents the domain-history command
var IntelDomainHistoryGetCmd = &cobra.Command{
	Use:   "domain-history",
	Short: "Intel.DomainHistory.Get for Intel.DomainHistory.Get",
	Long:  "account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Intel.DomainHistory.Get(ctx, intel.DomainHistoryGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response)
	},
}
