package service

type JwtService struct {
	JWT_CERTIFICATE string
	storage         JwtStorage
}

func NewJwtService(storage JwtStorage, jwtCertificate string) *JwtService {
	return &JwtService{
		JWT_CERTIFICATE: jwtCertificate,
		storage:         storage,
	}
}

type JwtStorage interface {
}
