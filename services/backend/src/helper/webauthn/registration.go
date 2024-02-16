package webauthn

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

func (webN *WebAuthN) BeginRegistration(user *WebAuthNUser) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	options, session, err := webN.webAuthn.BeginRegistration(user)
	if err != nil {
		return nil, nil, err
	}
	return options, session, nil
}

func (webN *WebAuthN) FinishRegistration() {
	// _, err := webN.webAuthn.FinishRegistration(user, session, r)
	// if err != nil {
	// 	// Handle Error and return.

	// 	return
	// }

}
