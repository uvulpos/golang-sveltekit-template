package service

func (s *GeneralSvc) Ping() error {
	return s.storage.PingDatabase()
}
