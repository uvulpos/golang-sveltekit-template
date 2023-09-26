package service

func (h *UserSvc) GetUserByUUID(uuid string) (*UserWithPermission, error) {
	return h.storage.GetUserByUUID(uuid)
}
