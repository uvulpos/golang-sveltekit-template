package service

func (s *AuthService) Logout() (string, error) {
	return s.authentikOauthLogoutEP, nil
}
