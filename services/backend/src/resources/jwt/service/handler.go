package service

type JwtService struct {
	JwtTokenValidInMinutes  int
	RefreshTokenValidInDays int
}

func NewJwtService(jwtCertificate string) *JwtService {
	return &JwtService{
		JwtTokenValidInMinutes:  10,
		RefreshTokenValidInDays: 30,
	}
}
