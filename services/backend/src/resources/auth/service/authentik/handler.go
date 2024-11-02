package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"golang.org/x/oauth2"

	userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
)

type AuthService struct {
	authentikConfig          *oauth2.Config
	authentikOauthUserInfoEP string
	authentikOauthLogoutEP   string

	userSvc *userService.UserService
}

func NewAuthService(userSvc *userService.UserService) *AuthService {
	authentikConfig := &oauth2.Config{
		ClientID:     configuration.AUTHORIZATION_OAUTH_AUTHENTIK_KEY,
		ClientSecret: configuration.AUTHORIZATION_OAUTH_AUTHENTIK_SECRET,

		RedirectURL: configuration.AUTHORIZATION_OAUTH_AUTHENTIK_CALLBACK_URL,
		Scopes:      configuration.AUTHORIZATION_OAUTH_AUTHENTIK_SCOPES,

		Endpoint: oauth2.Endpoint{
			AuthURL:  configuration.AUTHORIZATION_OAUTH_AUTHENTIK_AUTHORIZATION_URL,
			TokenURL: configuration.AUTHORIZATION_OAUTH_AUTHENTIK_TOKEN_URL,

			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	return &AuthService{
		authentikConfig:          authentikConfig,
		authentikOauthUserInfoEP: configuration.AUTHORIZATION_OAUTH_AUTHENTIK_USERINFO_URL,
		authentikOauthLogoutEP:   configuration.AUTHORIZATION_OAUTH_AUTHENTIK_LOGOUT_URL,

		userSvc: userSvc,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface)
	GetUserIDByLogin(provider string, providerID string) (string, customerrors.ErrorInterface)
	CreateUser(tx *sqlx.Tx, displayName string, username string, email string, emailVerified bool) (string, customerrors.ErrorInterface)
	CreateUserLoginIdentity(tx *sqlx.Tx, createdUserID string, authProvider string, authProviderID string) customerrors.ErrorInterface
}
