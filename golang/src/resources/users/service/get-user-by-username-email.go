package service

func (h *UserSvc) GetProfileByUsernameOrEmail(accountidentifier string) (*UserWithPermission, error) {
	return h.storage.GetUserByUsernameOrEmail(accountidentifier)
}
