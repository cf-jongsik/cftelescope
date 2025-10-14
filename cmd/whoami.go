/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Get Cloudflare account information",
	Long:  `Get Cloudflare account information`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		accounts := cf.Accounts.ListAutoPaging(ctx, accounts.AccountListParams{})
		prevAccountID := ""
		for accounts.Next() {
			//! there is bug in SDK that accounts.Next() will go unlimited
			if prevAccountID == accounts.Current().ID {
				break
			}
			prevAccountID = accounts.Current().ID
			fmt.Println(accounts.Current().JSON.RawJSON())
		}
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
