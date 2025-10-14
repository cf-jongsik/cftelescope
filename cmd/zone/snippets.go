package zone

import (
	"cftelescope/cmd/zone/snippets"

	"github.com/spf13/cobra"
)

// SnippetsCmd represents the snippets command
var SnippetsCmd = &cobra.Command{
	Use:   "snippets",
	Short: "Snippets commands",
	Long:  "Snippets commands",
}

func init() {
	SnippetsCmd.AddCommand(snippets.SnippetsGetCmd)
	SnippetsCmd.AddCommand(snippets.SnippetsListCmd)
}
