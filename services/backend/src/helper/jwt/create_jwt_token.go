package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	models "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt/models"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func CreateJWT(authProvider provider.AuthProvider, jwtData *models.JwtDataModel) (string, customerrors.ErrorInterface) {
	validityMinutes := time.Duration(configuration.JWT_TOKEN_VALIDITY_IN_MINUTES)

	claims := jwt.MapClaims{
		"auth-provider": authProvider,
		"user-uuid":     jwtData.UserUuid,
		"session-uuid":  jwtData.SessionID,
		"scopes":        jwtData.Scopes,
		"exp":           time.Now().Add(time.Minute * validityMinutes).Unix(),
		"authorized":    true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configuration.JWT_SIGNING_KEY))
	if err != nil {
		return "", customerrors.NewInternalServerError(err, jwtData.UserUuid, fmt.Sprintf("Could not sign jwt on creation %v", claims))
	}

	return tokenString, nil
}
