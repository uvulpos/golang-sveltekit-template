package service

func (s *GeneralSvc) SystemHealthCheck() error {

	// ping database
	if err := s.PingDatabase(); err != nil {
		return err
	}

	// ping kafka
	// ping s3 storage
	// ping ...

	return nil
}
