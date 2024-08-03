package service

import (
	"crypto/ecdsa"
	"crypto/x509"
)

type X509Certificat struct {
	content     []byte
	certificate *x509.Certificate
}

type X509PrivateKey struct {
	content    []byte
	privateKey *ecdsa.PrivateKey
}

func NewX509Certificat(content []byte, certificate *x509.Certificate) *X509Certificat {
	return &X509Certificat{
		content,
		certificate,
	}
}

func (c X509Certificat) ToByte() []byte {
	return c.content
}

func (c X509Certificat) ToString() string {
	return string(c.content)
}

func (c X509Certificat) GetCertificate() *x509.Certificate {
	return c.certificate
}

func NewX509PrivateKey(content []byte, privateKey *ecdsa.PrivateKey) *X509PrivateKey {
	return &X509PrivateKey{
		content,
		privateKey,
	}
}

func (c X509PrivateKey) ToByte() []byte {
	return c.content
}

func (c X509PrivateKey) ToString() string {
	return string(c.content)
}

func (c X509PrivateKey) GetPrivateKey() *ecdsa.PrivateKey {
	return c.privateKey
}
