package storage

func (s *GeneralStore) PingDatabase() error {
	return s.dbstore.Ping()
}
