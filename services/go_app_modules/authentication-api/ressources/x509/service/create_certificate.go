package service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	mathRand "math/rand"
	"net"
	"time"
)

type CustomX509CertificatModel struct {
	Name             string
	Organization     string
	OrganizationUnit string

	Address    string
	City       string
	PostalCode string
	State      string
	Country    string
}

type CustomSslCertificatModel struct {
	DNSNames    []string
	IPAddresses []net.IP
}

func CreateCertificate(parentCertificate *x509.Certificate, parentPrivateKey *ecdsa.PrivateKey, x509CertificateModel CustomX509CertificatModel, sslCertificateModel *CustomSslCertificatModel) (*X509Certificat, *X509PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Create a template for the certificate
	template := x509.Certificate{
		SerialNumber: big.NewInt(mathRand.Int63()),
		Subject: pkix.Name{
			Names:      []pkix.AttributeTypeAndValue{},
			ExtraNames: []pkix.AttributeTypeAndValue{},

			CommonName:         x509CertificateModel.Name,
			Organization:       []string{x509CertificateModel.Organization},
			OrganizationalUnit: []string{x509CertificateModel.OrganizationUnit},

			StreetAddress: []string{x509CertificateModel.Address},
			Locality:      []string{x509CertificateModel.City},
			PostalCode:    []string{x509CertificateModel.PostalCode},
			Province:      []string{x509CertificateModel.State},
			Country:       []string{x509CertificateModel.Country},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(0, 3, 0), // 1,5 year validity

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	if parentCertificate == nil {
		template.IsCA = true
	}

	if sslCertificateModel != nil {
		template.DNSNames = sslCertificateModel.DNSNames
		template.IPAddresses = sslCertificateModel.IPAddresses
	}

	if parentCertificate == nil {
		parentCertificate = &template
	}

	signingKey := privateKey
	if parentPrivateKey != nil {
		signingKey = parentPrivateKey
	}

	// Create a self-signed certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, parentCertificate, &privateKey.PublicKey, signingKey)
	if err != nil {
		return nil, nil, err
	}

	certificateBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})

	privBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}

	privateKeyBytes := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})

	certificatModel := NewX509Certificat(certificateBytes, &template)
	privateKeyModel := NewX509PrivateKey(privateKeyBytes, privateKey)

	err = ValidateCertificateAndKey(certificatModel.ToString(), privateKeyModel.ToString())
	if err != nil {
		return nil, nil, err
	}

	return certificatModel, privateKeyModel, nil
}
