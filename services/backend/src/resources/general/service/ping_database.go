package service

func (s *GeneralSvc) PingDatabase() error {
	return s.storage.PingDatabase()
}
