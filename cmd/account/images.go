package account

import (
	"cftelescope/cmd/account/images"

	"github.com/spf13/cobra"
)

var ImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Images commands",
	Long:  "images commands for account",
}

func init() {
	ImagesCmd.AddCommand(images.ImagesV1GetCmd)
	ImagesCmd.AddCommand(images.ImagesV2ListCmd)
}
