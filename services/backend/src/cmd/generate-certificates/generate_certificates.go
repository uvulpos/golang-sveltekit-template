package generatecertificates

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	certificateService "github.com/uvulpos/go-svelte/authentication-api/ressources/x509/service"
	"github.com/uvulpos/go-svelte/src/configuration"
	"github.com/uvulpos/go-svelte/src/helper/branding"
)

var GenerateCertificateCmd = &cobra.Command{
	Use:   "generate-certificates",
	Short: "📝 generate required certificates for secure communication",
	Long:  `📝 generate required certificates for secure communication`,
	Run: func(cmd *cobra.Command, args []string) {

		branding.PrintBranding()

		config := certificateService.CustomX509CertificatModel{
			Name:             configuration.CERTIFICATE_IDENTITY_NAME,
			Organization:     configuration.CERTIFICATE_IDENTITY_ORGANIZATION,
			OrganizationUnit: configuration.CERTIFICATE_IDENTITY_ORGANIZATION_UNIT,

			Address:    configuration.CERTIFICATE_IDENTITY_ORGANIZATION_ADRESS,
			City:       configuration.CERTIFICATE_IDENTITY_ORGANIZATION_CITY,
			PostalCode: configuration.CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_CODE,
			State:      configuration.CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_STATE,
			Country:    configuration.CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_COUNTRY,
		}

		sslConfig := &certificateService.CustomSslCertificatModel{
			IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		}

		// jwtCertificate, jwtPrivateKey, jwtErr := certificateService.CreateCertificate(nil, nil, config, nil)
		// if jwtErr != nil {
		// 	panic(jwtErr)
		// }

		caCertificate, caPrivateKey, caErr := certificateService.CreateCertificate(nil, nil, config, sslConfig)
		if caErr != nil {
			panic(caErr)
		}

		config.Name += "Sub Domäne"
		config.OrganizationUnit = "Drecksack Appartment"

		sslCertificate, sslPrivateKey, sslErr := certificateService.CreateCertificate(
			caCertificate.GetCertificate(),
			caPrivateKey.GetPrivateKey(),
			config, sslConfig,
		)
		if sslErr != nil {
			panic(sslErr)
		}

		fullCert := fmt.Sprintf("%s\n%s", sslCertificate.ToString(), caCertificate.ToString())

		fmt.Println()
		fmt.Println()
		fmt.Println("SSL-CERTIFICATE")
		fmt.Println(fullCert)
		fmt.Println()
		fmt.Println()
		fmt.Println("SSL-PRIVATE-KEY")
		fmt.Println(sslPrivateKey.ToString())
		fmt.Println()
		fmt.Println()

	},
}
