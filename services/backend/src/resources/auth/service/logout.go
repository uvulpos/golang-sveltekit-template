package service

func (s *AuthService) Logout() (string, error) {
	return s.auth.Logout()
}
