package models

type UserSelfInformation struct {
	Permissions []string
}

func NewUserSelfInformation(permissions []string) *UserSelfInformation {
	return &UserSelfInformation{
		Permissions: permissions,
	}
}
