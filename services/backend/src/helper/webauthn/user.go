package webauthn

import "github.com/go-webauthn/webauthn/webauthn"

type WebAuthNUser struct {
	ID          string
	Name        string
	DisplayName string
	Icon        string
	Credentials []webauthn.Credential
}

func CreateWebAuthNUser(
	ID,
	Name,
	DisplayName string,
	Credentials []webauthn.Credential,
	CredentialsSignIn []webauthn.Credential,
) *WebAuthNUser {
	return &WebAuthNUser{
		ID,
		Name,
		DisplayName,
		"",
		Credentials,
	}
}

func CreateWebAuthNUserWithOutCertificates(
	ID,
	Name,
	DisplayName string,
) *WebAuthNUser {
	return &WebAuthNUser{
		ID,
		Name,
		DisplayName,
		"",
		[]webauthn.Credential{},
	}
}

func (u *WebAuthNUser) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u *WebAuthNUser) WebAuthnName() string {
	return u.Name
}

func (u *WebAuthNUser) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *WebAuthNUser) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

// WebAuthnIcon is a deprecated option.
// Deprecated: this has been removed from the specification recommendation. Suggest a blank string.
func (u *WebAuthNUser) WebAuthnIcon() string {
	return ""
}
