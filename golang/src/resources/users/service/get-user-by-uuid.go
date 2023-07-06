package service

func (h *UserSvc) GetUserByUUID(uuid string) {
	h.storage.GetUserByUUID(uuid)
}
