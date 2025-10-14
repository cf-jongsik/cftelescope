package ssl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/ssl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// typed response for Certificate Pack
type SSLCertificatePack struct {
	CertificateAuthority string           `json:"certificate_authority"`
	Certificates         []SSLCertificate `json:"certificates"`
	CreatedOn            *time.Time       `json:"created_on,omitempty"`
	Hosts                []string         `json:"hosts,omitempty"`
	ID                   string           `json:"id"`
	PrimaryCertificate   string           `json:"primary_certificate,omitempty"`
	Status               string           `json:"status,omitempty"`
	Type                 string           `json:"type,omitempty"`
	ValidationMethod     string           `json:"validation_method,omitempty"`
	ValidityDays         *int             `json:"validity_days,omitempty"`
}

type SSLCertificate struct {
	BundleMethod string     `json:"bundle_method,omitempty"`
	ExpiresOn    *time.Time `json:"expires_on,omitempty"`
	Hosts        []string   `json:"hosts,omitempty"`
	ID           string     `json:"id"`
	Issuer       string     `json:"issuer,omitempty"`
	ModifiedOn   *time.Time `json:"modified_on,omitempty"`
	Priority     *string    `json:"priority,omitempty"`
	Signature    string     `json:"signature,omitempty"`
	Status       string     `json:"status,omitempty"`
	UploadedOn   *time.Time `json:"uploaded_on,omitempty"`
	ZoneID       string     `json:"zone_id,omitempty"`
}

// SslCertificatePackListCmd represents the verification command
var SslCertificatePackListCmd = &cobra.Command{
	Use:   "packs",
	Short: "SSL.CertificatePack.List for SSL.CertificatePack.List",
	Long:  "zone",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("cloudflare_api_token")
		zoneID := viper.GetString("cloudflare_zone_id")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		certs := cf.SSL.CertificatePacks.ListAutoPaging(ctx, ssl.CertificatePackListParams{ZoneID: cloudflare.F(zoneID), Status: cloudflare.F(ssl.CertificatePackListParamsStatusAll)})
		for certs.Next() {
			binary, err := json.Marshal(certs.Current())
			if err != nil {
				log.Println("marshal dynamic result:", err)
				continue
			}
			var pack SSLCertificatePack
			if err := json.Unmarshal(binary, &pack); err != nil {
				log.Println("unmarshal into SSLCertificatePack:", err)
				// fallback to printing raw
				fmt.Println(string(binary))
				continue
			}
			fmt.Println(pack.Hosts, pack.Type, pack.Status, pack.ID)
			for _, cert := range pack.Certificates {
				fmt.Println("\t", cert.Status, cert.ID, cert.ExpiresOn, cert.Issuer)
			}
		}
	},
}
