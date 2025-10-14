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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Override .env values with flag values if flags were explicitly provided
		overrideFlagIfChanged(cmd, "cloudflare_account_id")
		overrideFlagIfChanged(cmd, "cloudflare_api_token")
		overrideFlagIfChanged(cmd, "cloudflare_zone_id")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

// overrideFlagIfChanged checks if a flag was explicitly provided and overrides the viper value
func overrideFlagIfChanged(cmd *cobra.Command, flagName string) {
	flag := cmd.Flags().Lookup(flagName)
	if flag != nil && flag.Changed {
		viper.Set(flagName, flag.Value.String())
	}
}

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringP("cloudflare_account_id", "a", "", "Account ID")
	viper.BindPFlag("cloudflare_account_id", rootCmd.PersistentFlags().Lookup("cloudflare_account_id"))

	rootCmd.PersistentFlags().StringP("cloudflare_api_token", "t", "", "API Token")
	viper.BindPFlag("cloudflare_api_token", rootCmd.PersistentFlags().Lookup("cloudflare_api_token"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(".env file should be in the same directory as the binary\nPlease create a .env file with the following content:\nCLOUDFLARE_ACCOUNT_ID=your_account_id\nCLOUDFLARE_API_TOKEN=your_api_token")
		return
	}
}
