package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *JwtService) CreateJWT(jwtData *JwtDataModel) (string, customerrors.ErrorInterface) {
	claims := jwt.MapClaims{
		"user-uuid":    jwtData.UserUuid,
		"session-uuid": jwtData.SessionID,
		"exp":          time.Now().Add(time.Minute * 10).Unix(),
		"authorized":   true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.JWT_CERTIFICATE))
	if err != nil {
		return "", customerrors.NewInternalServerError(err, jwtData.UserUuid, fmt.Sprintf("Could not sign jwt on creation %v", claims))
	}

	return tokenString, nil
}
