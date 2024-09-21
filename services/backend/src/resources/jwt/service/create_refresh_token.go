package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *JwtService) CreateRefreshToken(sessionID string) (string, customerrors.ErrorInterface) {
	claims := jwt.MapClaims{
		"session-uuid": sessionID,
		"exp":          time.Now().Add(time.Minute * time.Duration(configuration.REFRESH_TOKEN_VALIDITY_IN_DAYS)).Unix(),
		"authorized":   true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configuration.REFRESH_TOKEN_SIGNING_KEY))
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", fmt.Sprintf("Could not sign refresh token on creation %v", claims))
	}

	return tokenString, nil
}
