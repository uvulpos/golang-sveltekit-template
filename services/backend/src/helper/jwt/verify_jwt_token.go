package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	models "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt/models"
)

func VerifyJWT(jwtString string) (*models.JwtDataModel, error) {

	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		_, okECDSA := token.Method.(*jwt.SigningMethodECDSA)
		_, okHS256 := token.Method.(*jwt.SigningMethodHMAC)

		if !okECDSA && !okHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configuration.JWT_SIGNING_KEY), nil
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

	var permissionScopes []string
	if value, ok := claims["scopes"].([]string); ok {
		permissionScopes = value
	}

	data := models.NewJwtDataModel(
		claims["user-uuid"].(string),
		claims["session-uuid"].(string),
		permissionScopes,
	)

	return data, nil
}
