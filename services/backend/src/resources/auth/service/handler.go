package service

import (
	"github.com/go-sqlx/sqlx"
	authService "github.com/uvulpos/go-svelte/authentication-api/ressources/auth/service"
	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

type AuthService struct {
	storage AuthStorageInterface

	auth *authService.AuthService
	jwt  *jwtService.JwtService
}

func NewAuthService(storage AuthStorageInterface, authService *authService.AuthService, jwtService *jwtService.JwtService) *AuthService {
	return &AuthService{
		storage: storage,
		auth:    authService,
		jwt:     jwtService,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	StartLoginSession(tx *sqlx.Tx, loggedinUser string, useragentHash, ipaddr *string) (sessionID string, err customerrors.ErrorInterface)
}
