package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
	"golang.org/x/oauth2"
)

type AuthService struct {
	authentikConfig          *oauth2.Config
	authentikOauthUserInfoEP string
	authentikOauthLogoutEP   string

	storage AuthStorageInterface
}

func NewAuthService(storage AuthStorageInterface, OAuthKey, OAuthSecret, CallbackURL, AuthURL, AuthTokenURL, UserInfoURL, LogoutURL string, scope ...string) *AuthService {
	authentikConfig := &oauth2.Config{
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
		authentikConfig:          authentikConfig,
		authentikOauthUserInfoEP: UserInfoURL,
		authentikOauthLogoutEP:   LogoutURL,

		storage: storage,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface)
	GetUserIDByLogin(provider string, providerID string) (string, customerrors.ErrorInterface)
	CreateUser(tx *sqlx.Tx, displayName string, username string, email string, emailVerified bool) (string, customerrors.ErrorInterface)
	CreateUserLoginIdentity(tx *sqlx.Tx, createdUserID string, authProvider string, authProviderID string) customerrors.ErrorInterface
}
