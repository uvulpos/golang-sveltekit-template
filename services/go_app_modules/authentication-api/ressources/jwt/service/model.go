package service

type JwtDataModel struct {
	UserUuid string `json:"user-uuid"`
}

func NewJwtDataModel(uuid string) *JwtDataModel {
	return &JwtDataModel{
		UserUuid: uuid,
	}
}
