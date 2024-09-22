package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

func (svc *JwtService) VerifyRefreshToken(refreshTokenString string) (*RefreshDataModel, error) {
	return VerifyRefreshToken(refreshTokenString)
}

func VerifyRefreshToken(refreshTokenString string) (*RefreshDataModel, error) {

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		_, okECDSA := token.Method.(*jwt.SigningMethodECDSA)
		_, okHS256 := token.Method.(*jwt.SigningMethodHMAC)

		if !okECDSA && !okHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configuration.REFRESH_TOKEN_SIGNING_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims["authorized"] != true {
		return nil, fmt.Errorf("unauthorized")
	}

	var sessionUuid string
	if val, ok := claims["session-uuid"]; ok {
		sessionUuid = val.(string)
	}

	data := NewRefreshDataModel(sessionUuid)
	return data, nil
}
