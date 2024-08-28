package service

import "fmt"

func (s *AuthService) Logout() (string, error) {
	fmt.Println("Logout", s.authentikOauthLogoutEP)
	return s.authentikOauthLogoutEP, nil
}
