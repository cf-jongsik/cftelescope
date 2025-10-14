package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cftelescope",
	Short: "Cloudflare Inspector",
	Long:  `cftelescope is a CLI tool for inspect Cloudflare accounts/zones`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(".env file should be in the same directory as the binary\nPlease create a .env file with the following content:\nCLOUDFLARE_ACCOUNT_ID=your_account_id\nCLOUDFLARE_API_TOKEN=your_api_token")
		return
	}
}
