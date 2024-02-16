package webauthn

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
)

type WebAuthN struct {
	webAuthn *webauthn.WebAuthn
}

func CreateNewWebAuthN(DisplayName, Domain string, AllowedDomains ...string) *WebAuthN {

	if len(AllowedDomains) == 0 {
		AllowedDomains = []string{Domain}
	}

	wconfig := &webauthn.Config{
		RPDisplayName: DisplayName,    // Display Name for your site
		RPID:          Domain,         // Generally the FQDN for your site
		RPOrigins:     AllowedDomains, // The origin URLs allowed for WebAuthn requests
	}

	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	return &WebAuthN{
		webAuthn,
	}
}

// // get Configuration of WebAuthN Handler
// func (webN *WebAuthN) GetConfig() *webauthn.Config {
// 	return webN.webAuthn.Config
// }
