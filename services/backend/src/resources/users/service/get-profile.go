package service

func (h *UserSvc) GetProfileByUserUUID(uuid string) (*UserWithPermission, error) {
	return h.storage.GetUserByUUID(uuid)
}
