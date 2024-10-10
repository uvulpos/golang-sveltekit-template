package service

func (s *AuthService) CreateRedirect(state string) string {
	return s.authentikProviderSvc.CreateRedirect(state)
}
