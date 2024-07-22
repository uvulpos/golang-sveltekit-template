package service

import (
	"golang.org/x/oauth2"
)

type AuthService struct {
	config          *oauth2.Config
	oauthUserInfoEP string
	oauthLogoutEP   string
}

func NewAuthService(OAuthKey, OAuthSecret, CallbackURL, AuthURL, AuthTokenURL, UserInfoURL, LogoutURL string, scope ...string) *AuthService {
	config := &oauth2.Config{
		ClientID:     OAuthKey,
		ClientSecret: OAuthSecret,

		RedirectURL: CallbackURL,
		Scopes:      scope,

		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: AuthTokenURL,

			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	return &AuthService{
		config:          config,
		oauthUserInfoEP: UserInfoURL,
		oauthLogoutEP:   LogoutURL,
	}
}
