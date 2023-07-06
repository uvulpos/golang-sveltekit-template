package service

func (h *UserSvc) GetProfileByUserUUID(uuid string) {
	h.GetUserByUUID(uuid)
}
