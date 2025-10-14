package intel

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/intel"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// IntelDNSListCmd represents the d-n-s command
var IntelDNSListCmd = &cobra.Command{
	Use:   "dns [ipv4]",
	Short: "Intel.DNS.List for Intel.DNS.List",
	Long:  "account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		ipv4 := args[0]
		response := cf.Intel.DNS.ListAutoPaging(ctx, intel.DNSListParams{AccountID: cloudflare.F(accountID), IPV4: cloudflare.F(ipv4), StartEndParams: cloudflare.F(intel.DNSListParamsStartEndParams{Start: cloudflare.F(time.Now().AddDate(0, 0, -3))})})
		for response.Next() {
			fmt.Println(response.Current().JSON.RawJSON())
		}
	},
}
