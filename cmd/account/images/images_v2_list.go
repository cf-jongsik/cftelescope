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

// ImagesV2ListCmd represents the v2 command
var ImagesV2ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list image",
	Long:  "list image for the specified account",
	Run: func(cmd *cobra.Command, args []string) {
		accountID := viper.GetString("cloudflare_account_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.Images.V2.List(ctx, images.V2ListParams{AccountID: cloudflare.F(accountID)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(response.JSON.Images.Raw())
		for {
			token := response.ContinuationToken
			if token == "" {
				break
			}
			response, err = cf.Images.V2.List(ctx, images.V2ListParams{AccountID: cloudflare.F(accountID), ContinuationToken: cloudflare.F(token)})
			if err != nil {
				log.Fatalln(err)
				return
			}
			fmt.Println(response.JSON.Images.Raw())
		}
	},
}
