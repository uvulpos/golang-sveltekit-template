package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *JwtService) CreateJWT(jwtData *JwtDataModel) (string, customerrors.ErrorInterface) {
	claims := jwt.MapClaims{
		"auth-provider": "authentik",
		"user-uuid":     jwtData.UserUuid,
		"session-uuid":  jwtData.SessionID,
		"scopes":        jwtData.Scopes,
		"exp":           time.Now().Add(time.Minute * time.Duration(configuration.JWT_TOKEN_VALIDITY_IN_MINUTES)).Unix(),
		"authorized":    true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configuration.JWT_SIGNING_KEY))
	if err != nil {
		return "", customerrors.NewInternalServerError(err, jwtData.UserUuid, fmt.Sprintf("Could not sign jwt on creation %v", claims))
	}

	return tokenString, nil
}
