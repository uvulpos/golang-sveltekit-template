package generatecertificates

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	certificateService "github.com/uvulpos/go-svelte/authentication-api/ressources/x509/service"
	"github.com/uvulpos/go-svelte/src/configuration"
	"github.com/uvulpos/go-svelte/src/helper/branding"
)

const caCertString = `-----BEGIN CERTIFICATE-----
MIIC2jCCAoCgAwIBAgIIUtk0y8emX5owCgYIKoZIzj0EAwIwgb0xEDAOBgNVBAYT
B0dlcm1hbnkxEDAOBgNVBAgTB0JhdmFyaWExDzANBgNVBAcTBk11bmljaDEbMBkG
A1UECRMSRXhhbXBsZSBTdHJlZXQgMTJhMQ4wDAYDVQQREwUxMjM0NTEXMBUGA1UE
ChMOVGhlIElubm92YXRvcnMxHTAbBgNVBAsTFFRlY2huaWNhbCBJbm5vdmF0aW9u
MSEwHwYDVQQDExh1VnVscG9zIC0gTXkgQXBwbGljYXRpb24wHhcNMjQwNzE0MjMz
MjE4WhcNMjYwMTE1MDAzMjE4WjCBvTEQMA4GA1UEBhMHR2VybWFueTEQMA4GA1UE
CBMHQmF2YXJpYTEPMA0GA1UEBxMGTXVuaWNoMRswGQYDVQQJExJFeGFtcGxlIFN0
cmVldCAxMmExDjAMBgNVBBETBTEyMzQ1MRcwFQYDVQQKEw5UaGUgSW5ub3ZhdG9y
czEdMBsGA1UECxMUVGVjaG5pY2FsIElubm92YXRpb24xITAfBgNVBAMTGHVWdWxw
b3MgLSBNeSBBcHBsaWNhdGlvbjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFZ7
6vDdKVczikplpMGqGXh8eRqbuNWZwijslXZHqj+i0SE1k9sSZ350RNGeR1CNfjAf
mgjK5w6VK5xQFxrn84mjaDBmMA4GA1UdDwEB/wQEAwIFoDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBT4OtuXbLYQTUZMDb5Z
F+qSC+aKqTAPBgNVHREECDAGhwR/AAABMAoGCCqGSM49BAMCA0gAMEUCIDOl15Td
iLgVAXdDC290BvXhSCBtNSJJpAvede6263UiAiEA824FMtBMdpdLHVEyT5bWHw5m
NuJYLyfyrkqCVo4u778=
-----END CERTIFICATE-----`

const caCertPrivateKeyString = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIS0r90LxEkHEfo7HWWqLkT/AcYwHSXQoEg3ZnGY6AsOoAoGCCqGSM49
AwEHoUQDQgAEVnvq8N0pVzOKSmWkwaoZeHx5Gpu41ZnCKOyVdkeqP6LRITWT2xJn
fnRE0Z5HUI1+MB+aCMrnDpUrnFAXGufziQ==
-----END EC PRIVATE KEY-----`

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

		// caCertificate, caPrivateKey, caErr := certificateService.CreateCertificate(nil, nil, config, sslConfig)
		// if caErr != nil {
		// 	panic(caErr)
		// }

		caCert, _ := certificateService.ConvertBytesToX509Certificate([]byte(caCertString))

		caPrivKey, caPrivKeyErr := certificateService.ConvertBytesToX509PrivateKey([]byte(caCertPrivateKeyString))
		if caPrivKeyErr != nil {
			panic(caPrivKeyErr)
		}

		config.Name += "Sub Domäne"
		config.OrganizationUnit = "Drecksack Appartment"

		sslCert, sslCertPrivateKey, sslErr := certificateService.CreateCertificate(
			caCert,
			caPrivKey,
			config, sslConfig,
		)
		if sslErr != nil {
			panic(sslErr)
		}

		fullCert := fmt.Sprintf("%s\n%s", sslCert.ToString(), caCertString)

		// fmt.Println()
		// fmt.Println()
		// fmt.Println("JWT-CERTIFICATE")
		// fmt.Println(jwtCertificate.ToString())
		// fmt.Println()
		// fmt.Println()
		// fmt.Println("JWT-PRIVATE-KEY")
		// fmt.Println(jwtPrivateKey.ToString())
		// fmt.Println()
		// fmt.Println()
		fmt.Println("------------------------------------------------------------------------")
		fmt.Println()
		fmt.Println()
		fmt.Println("SSL-CERTIFICATE")
		fmt.Println(fullCert)
		fmt.Println()
		fmt.Println()
		fmt.Println("SSL-PRIVATE-KEY")
		fmt.Println(sslCertPrivateKey.ToString())
		fmt.Println()
		fmt.Println()

	},
}
