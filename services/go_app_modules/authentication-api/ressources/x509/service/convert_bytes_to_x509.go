package service

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ConvertBytesToX509Certificate(certificate []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certificate)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("WIIIIIILLDDD")
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("BOOODENLOS!")
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}

	return cert, nil
}

func ConvertBytesToX509PrivateKey(privateKey []byte) (*ecdsa.PrivateKey, error) {
	// Parse the private key
	block, _ := pem.Decode(privateKey)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		fmt.Println(block.Type)
		return nil, fmt.Errorf("failed to parse private key PEM")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	return key, nil
}
