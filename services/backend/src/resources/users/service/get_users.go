package service

func (h *UserSvc) GetUsers() ([]*User, error) {
	return h.storage.GetUsers(nil)
}
