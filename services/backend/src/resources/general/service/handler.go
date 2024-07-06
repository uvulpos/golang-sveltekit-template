package service

type GeneralSvc struct {
	storage GeneralStorage
}

func NewGeneralSvc(storage GeneralStorage) *GeneralSvc {
	return &GeneralSvc{
		storage,
	}
}

type GeneralStorage interface {
	PingDatabase() error
}
