/*
Copyright © 2025 jongsik <jongsik@cloudflare.com>
*/
package zone

import (
	"cftelescope/cmd/zone/argo"

	"github.com/spf13/cobra"
)

// ArgoCmd represents the argo command
var ArgoCmd = &cobra.Command{
	Use:   "argo [cloudflare_zone_id]",
	Short: "Argo commands for a zone",
	Long:  `Argo commands for the specified zone`,
}

func init() {
	ArgoCmd.AddCommand(argo.ArgoSmartRoutingGetCmd)
	ArgoCmd.AddCommand(argo.ArgoTieredCachingGetCmd)
}
