package images

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/images"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ImagesV1GetCmd represents the v1 command
var ImagesV1GetCmd = &cobra.Command{
	Use:   "get [imageID]",
	Short: "get image",
	Long:  "get image for the specified account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		imageID := args[0]
		response, err := cf.Images.V1.Get(ctx, imageID, images.V1GetParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.RawJSON())
	},
}
