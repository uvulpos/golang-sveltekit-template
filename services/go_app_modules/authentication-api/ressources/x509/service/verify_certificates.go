package service

import (
	"fmt"
	"time"
)

func ValidateCertificateAndKey(certPEM, keyPEM string) error {

	cert, err := ConvertBytesToX509Certificate([]byte(certPEM))
	if err != nil {
		return err
	}

	key, err := ConvertBytesToX509PrivateKey([]byte(keyPEM))
	if err != nil {
		return err
	}

	// Check that the public key in the certificate matches the private key
	if !key.PublicKey.Equal(cert.PublicKey) {
		return fmt.Errorf("public key in certificate does not match private key")
	}

	// Additional checks like validating the certificate chain can be added here
	// Currently, we are only checking the validity period and the public key
	if cert.NotBefore.After(time.Now()) || cert.NotAfter.Before(time.Now()) {
		return fmt.Errorf("certificate is not within its validity period")
	}

	return nil
}
