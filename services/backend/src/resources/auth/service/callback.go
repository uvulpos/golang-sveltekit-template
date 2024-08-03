package service

import (
	"fmt"

	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
)

func (s *AuthService) CallbackFunction(authCode, state, oauthUserinfoURL string) (string, error) {
	oauthUser, oauthUserErr := s.auth.CallbackFunction(authCode, state, oauthUserinfoURL)
	if oauthUserErr != nil {
		fmt.Println("CALLBACK ERROR")
		return "", oauthUserErr
	}

	jwt, jwtErr := s.jwt.CreateJWT(jwtService.JwtDataModel{
		UserUuid: oauthUser.ID,
	})

	if jwtErr != nil {
		fmt.Println("JWT ERROR")
		return "", jwtErr
	}

	return jwt, nil
}
