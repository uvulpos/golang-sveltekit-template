package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *JwtService) CreateJWT(jwtData JwtDataModel) (string, error) {
	claims := jwt.MapClaims{
		"user-uuid":  jwtData.UserUuid,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		"authorized": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.JWT_CERTIFICATE))
	if err != nil {
		fmt.Println("JWT SIGNING ERROR")
		return "", err
	}

	return tokenString, nil
}
