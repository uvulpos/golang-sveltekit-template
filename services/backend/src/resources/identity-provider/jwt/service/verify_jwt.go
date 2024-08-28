package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func (svc *JwtService) VerifyJWT(jwtString string) (*JwtDataModel, error) {
	return VerifyJWT(jwtString, svc.JWT_CERTIFICATE)
}

func VerifyJWT(jwtString, signingKey string) (*JwtDataModel, error) {

	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		_, okECDSA := token.Method.(*jwt.SigningMethodECDSA)
		_, okHS256 := token.Method.(*jwt.SigningMethodHMAC)

		if !okECDSA && !okHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	fmt.Println(claims)

	if claims["authorized"] != true {
		return nil, fmt.Errorf("unauthorized")
	}

	data := NewJwtDataModel(
		claims["user-uuid"].(string),
		claims["session-uuid"].(string),
	)

	return data, nil
}
