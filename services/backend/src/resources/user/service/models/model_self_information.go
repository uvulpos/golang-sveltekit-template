package models

type UserSelfInformationModel struct {
	User        *UserModel
	Permissions []string
}

func NewUserSelfInformationModel(user *UserModel, permissions []string) *UserSelfInformationModel {
	return &UserSelfInformationModel{
		User:        user,
		Permissions: permissions,
	}
}
