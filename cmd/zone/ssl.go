/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package zone

import (
	sslPkg "cftelescope/cmd/zone/ssl"

	"github.com/spf13/cobra"
)

// SslCmd represents the ssl command
var SslCmd = &cobra.Command{
	Use:   "ssl",
	Short: "ssl commands",
	Long:  `ssl commands`,
}

func init() {
	SslCmd.AddCommand(sslPkg.SslCertificatePackListCmd)
	SslCmd.AddCommand(sslPkg.SslVerificationGetCmd)
	SslCmd.AddCommand(sslPkg.SslUniversalGetCmd)
	SslCmd.AddCommand(sslPkg.AcmTotalTLSGetCmd)
}
