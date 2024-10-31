package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt"
	authentikProviderService "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/authentik"
	userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
)

type AuthService struct {
	storage AuthStorageInterface

	jwt     *jwtService.JwtService
	userSvc *userService.UserService

	authentikProviderSvc *authentikProviderService.AuthService
}

func NewAuthService(storage AuthStorageInterface, jwtService *jwtService.JwtService, userService *userService.UserService) *AuthService {
	authentikProviderSvc := authentikProviderService.NewAuthService(userService)

	return &AuthService{
		storage: storage,
		jwt:     jwtService,
		userSvc: userService,

		authentikProviderSvc: authentikProviderSvc,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	StartLoginSession(tx *sqlx.Tx, loggedinUser string, useragentHash, ipaddr *string) (sessionID string, err customerrors.ErrorInterface)
}
