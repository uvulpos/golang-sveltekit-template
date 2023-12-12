package service

func (h *UserSvc) GetAuthenticatedUserByCredentials(username, password string) (*UserWithPermission, error) {
	return h.storage.GetUserByCredentials(username, password)
}
