package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func CreateRefreshToken(authProvider provider.AuthProvider, sessionID string) (string, customerrors.ErrorInterface) {
	validityDays := time.Duration(configuration.REFRESH_TOKEN_VALIDITY_IN_DAYS) * 24

	claims := jwt.MapClaims{
		"auth-provider": authProvider,
		"session-uuid":  sessionID,
		"exp":           time.Now().Add(time.Hour * validityDays).Unix(),
		"authorized":    true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configuration.REFRESH_TOKEN_SIGNING_KEY))
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", fmt.Sprintf("Could not sign refresh token on creation %v", claims))
	}

	return tokenString, nil
}
