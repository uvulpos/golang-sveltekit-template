package service

import "fmt"

func (s *AuthService) Logout() (string, error) {
	fmt.Println("Logout", s.oauthLogoutEP)
	return s.oauthLogoutEP, nil
}
