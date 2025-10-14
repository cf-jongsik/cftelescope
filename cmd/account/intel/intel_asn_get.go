package intel

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/intel"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Source struct {
	Pointer string `json:"pointer,omitempty"`
}

type ErrorMessage struct {
	Code             int     `json:"code"`
	Message          string  `json:"message"`
	DocumentationURL string  `json:"documentation_url,omitempty"`
	Source           *Source `json:"source,omitempty"`
}

type ASN struct {
	ASN         int    `json:"asn"`
	Description string `json:"description"`
	Country     string `json:"country"`
	Type        string `json:"type"`
}

type Response struct {
	Errors   []ErrorMessage `json:"errors"`
	Messages []ErrorMessage `json:"messages"`
	Success  bool           `json:"success"`
	Result   *ASN           `json:"result,omitempty"`
}

// IntelASNGetCmd represents the a-s-n command
var IntelASNGetCmd = &cobra.Command{
	Use:   "asn [asn]",
	Short: "Intel.ASN.Get for Intel.ASN.Get",
	Long:  "account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		// convert args[0] as int 64
		asnStr := args[0]
		asn, err := strconv.ParseInt(asnStr, 10, 64)
		if err != nil {
			log.Fatalln(err)
			return
		}
		response, err := cf.Intel.ASN.Get(ctx, asn, intel.ASNGetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println(response)
	},
}
