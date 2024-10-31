package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	authentikProviderService "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/authentik"
	userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
)

type AuthService struct {
	storage AuthStorageInterface

	userSvc *userService.UserService

	authentikProviderSvc *authentikProviderService.AuthService
}

func NewAuthService(storage AuthStorageInterface, userService *userService.UserService) *AuthService {
	authentikProviderSvc := authentikProviderService.NewAuthService(userService)

	return &AuthService{
		storage: storage,
		userSvc: userService,

		authentikProviderSvc: authentikProviderSvc,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	StartLoginSession(tx *sqlx.Tx, loggedinUser string, useragentHash, ipaddr *string) (sessionID string, err customerrors.ErrorInterface)
}
