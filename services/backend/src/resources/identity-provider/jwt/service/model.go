package service

type JwtDataModel struct {
	UserUuid  string `json:"user-uuid"`
	SessionID string `json:"session-uuid"`
}

func NewJwtDataModel(uuid, sessionID string) *JwtDataModel {
	return &JwtDataModel{
		UserUuid:  uuid,
		SessionID: sessionID,
	}
}
