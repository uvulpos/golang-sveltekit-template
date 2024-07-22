package service

import "golang.org/x/oauth2"

func (s *AuthService) CreateRedirect(state string) string {
	return s.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}
