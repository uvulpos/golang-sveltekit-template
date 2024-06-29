package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uvulpos/go-svelte/src/configuration"
)

type AppJWTClaims struct {
	Username    string   `json:"username"`
	Uuid        string   `json:"user-uuid"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func VerifyJWToken(jwtToken string) (bool, *AppJWTClaims, error) {

	if jwtToken == "" {
		return false, nil, errors.New("no jwt set")
	}

	token, tokenErr := jwt.ParseWithClaims(jwtToken, &AppJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configuration.Configuration.Session.JwtSecret), nil
	})
	if tokenErr != nil {
		fmt.Println("Verify JWT", tokenErr)
		return false, nil, errors.New("could not verify jwt token")
	}

	jwtClaims, jwtClaimsOk := token.Claims.(*AppJWTClaims)
	if !jwtClaimsOk {
		return false, nil, errors.New("could not verify jwt token")
	}

	return token.Valid, jwtClaims, nil
}
