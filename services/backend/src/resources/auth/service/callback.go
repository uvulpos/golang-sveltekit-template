package service

import (
	"fmt"

	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
)

func (s *AuthService) CallbackFunction(authCode, state, oauthUserinfoURL string) (string, error) {
	fmt.Println("service")
	loggedinUser, loggedinUserErr := s.auth.AuthentikCallbackFunction(authCode, state, oauthUserinfoURL)
	if loggedinUserErr != nil {
		return "", loggedinUserErr
	}

	jwt, jwtErr := s.jwt.CreateJWT(jwtService.NewJwtDataModel(loggedinUser))
	if jwtErr != nil {
		fmt.Println("JWT ERROR")
		return "", jwtErr
	}

	return jwt, nil
}
