/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package account

import (
	"cftelescope/cmd/account/pages"

	"github.com/spf13/cobra"
)

var PagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "pages commands for account",
	Long:  `pages commands for the specified account`,
}

func init() {
	PagesCmd.AddCommand(pages.PagesProjectListCmd)
	PagesCmd.AddCommand(pages.PagesProjectGetCmd)
}
