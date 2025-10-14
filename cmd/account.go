/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"cftelescope/cmd/account"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountCmd represents the account command group
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage Cloudflare accounts",
	Long:  `Manage Cloudflare accounts - list accounts`,
}

var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts",
	Long:  `List all accounts accessible with the API token`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		accounts := cf.Accounts.ListAutoPaging(ctx, accounts.AccountListParams{})
		prev_account := ""
		for accounts.Next() {
			//! there is glitch in the SDK, if you dont break, it will print the same account forever
			//! this is to prevent that
			if prev_account == accounts.Current().JSON.RawJSON() {
				break
			}

			prev_account = accounts.Current().JSON.RawJSON()
			fmt.Println(prev_account)
		}
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(accountListCmd)

	// Add account-specific subcommands
	accountCmd.AddCommand(account.ListListCmd)
	accountCmd.AddCommand(account.PagesCmd)
	accountCmd.AddCommand(account.WorkersCmd)
	accountCmd.AddCommand(account.AccountsAccountProfileGetCmd)
	accountCmd.AddCommand(account.AlertingHistoryListCmd)
	accountCmd.AddCommand(account.CallsCmd)
	accountCmd.AddCommand(account.D1Cmd)
	accountCmd.AddCommand(account.ImagesCmd)
	accountCmd.AddCommand(account.IntelCmd)
	accountCmd.AddCommand(account.LogsControlCmbConfigGetCmd)
	accountCmd.AddCommand(account.RUMsCmd)
	accountCmd.AddCommand(account.StreamCmd)

}
