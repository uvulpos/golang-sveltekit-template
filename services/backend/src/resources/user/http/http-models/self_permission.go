package httpModels

type SelfPermissionsModel struct {
	Permissions []string `json:"permissions"`
}

func NewSelfPermissionsModel(permissions []string) *SelfPermissionsModel {
	return &SelfPermissionsModel{
		Permissions: permissions,
	}
}
