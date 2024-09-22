package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	authService "github.com/uvulpos/golang-sveltekit-template/src/resources/identity-provider/auth/service"
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
	userService "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service"
)

type AuthService struct {
	storage AuthStorageInterface

	auth    *authService.AuthService
	jwt     *jwtService.JwtService
	userSvc *userService.UserService
}

func NewAuthService(storage AuthStorageInterface, authService *authService.AuthService, jwtService *jwtService.JwtService, userService *userService.UserService) *AuthService {
	return &AuthService{
		storage: storage,
		auth:    authService,
		jwt:     jwtService,
		userSvc: userService,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	StartLoginSession(tx *sqlx.Tx, loggedinUser string, useragentHash, ipaddr *string) (sessionID string, err customerrors.ErrorInterface)
}
