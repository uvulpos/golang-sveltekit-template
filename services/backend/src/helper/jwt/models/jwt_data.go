package service

type JwtDataModel struct {
	UserUuid  string   `json:"user-uuid"`
	SessionID string   `json:"session-uuid"`
	Scopes    []string `json:"scopes"`
}

func NewJwtDataModel(uuid, sessionID string, scopes []string) *JwtDataModel {
	return &JwtDataModel{
		UserUuid:  uuid,
		SessionID: sessionID,
		Scopes:    scopes,
	}
}

type RefreshDataModel struct {
	SessionID string `json:"session-uuid"`
}

func NewRefreshDataModel(sessionID string) *RefreshDataModel {
	return &RefreshDataModel{
		SessionID: sessionID,
	}
}
