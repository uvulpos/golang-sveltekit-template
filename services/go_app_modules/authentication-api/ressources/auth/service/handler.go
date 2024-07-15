package service

import (
	"golang.org/x/oauth2"
)

type AuthService struct {
	config *oauth2.Config
}

func NewAuthService(OAuthKey, OAuthSecret, CallbackURL string, scope ...string) *AuthService {
	config := &oauth2.Config{
		ClientID:     OAuthKey,
		ClientSecret: OAuthSecret,
		RedirectURL:  CallbackURL,
		Scopes:       scope,
		Endpoint: oauth2.Endpoint{
			AuthURL:       "https://example.com",
			TokenURL:      "",
			DeviceAuthURL: "",

			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	return &AuthService{
		config,
	}
}
