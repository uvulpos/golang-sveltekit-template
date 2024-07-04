package storage

func (s *GeneralStore) PingDatabase() error {
	return s.dbstore.DB.Ping()
}
