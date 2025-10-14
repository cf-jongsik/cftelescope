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

// SSLVerification represents the verification response

type SSLVerification struct {
	CertificateStatus  string            `json:"certificate_status,omitempty"`
	BrandCheck         *bool             `json:"brand_check,omitempty"`
	CertPackUUID       string            `json:"cert_pack_uuid,omitempty"`
	Signature          string            `json:"signature,omitempty"`
	ValidationMethod   string            `json:"validation_method,omitempty"`
	VerificationInfo   *VerificationInfo `json:"verification_info,omitempty"`
	VerificationStatus *bool             `json:"verification_status,omitempty"`
	VerificationType   string            `json:"verification_type,omitempty"`
}

type VerificationInfo struct {
	RecordName   string `json:"record_name,omitempty"`
	RecordTarget string `json:"record_target,omitempty"`
}

type SSLVerificationResponse struct {
	Result []SSLVerification `json:"result"`
}

// SslVerificationGetCmd represents the verification command
var SslVerificationGetCmd = &cobra.Command{
	Use:   "verification",
	Short: "SSL.Verification.Get for SSL.Verification.Get",
	Long:  "zone",
	Run: func(cmd *cobra.Command, args []string) {
		zoneID := viper.GetString("cloudflare_zone_id")
		apiToken := viper.GetString("cloudflare_api_token")
		cf := cloudflare.NewClient(option.WithAPIToken(apiToken))
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		response, err := cf.SSL.Verification.Get(ctx, ssl.VerificationGetParams{ZoneID: cloudflare.F(zoneID)})
		if err != nil {
			log.Fatalln(err)
			return
		}

		binary, err := json.Marshal(response)
		if err != nil {
			log.Println("marshal dynamic result:", err)
			fmt.Println(response)
			return
		}
		var verResp SSLVerificationResponse
		if err := json.Unmarshal(binary, &verResp); err != nil {
			log.Println("unmarshal into SSLVerificationResponse:", err)
			// fallback to printing raw
			fmt.Println(string(binary))
			return
		}
		for _, ver := range verResp.Result {
			fmt.Println(ver)
		}
	},
}
