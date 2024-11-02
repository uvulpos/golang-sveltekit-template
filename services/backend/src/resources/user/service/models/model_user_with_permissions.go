package models

type UserWithPermissionsModel struct {
	ID            string
	Username      string
	DisplayName   string
	Email         string
	EmailVerified bool
	Permissions   []string
}

func NewUserWithPermissionsModel(user *UserModel, permissions []string) *UserWithPermissionsModel {
	return &UserWithPermissionsModel{
		ID:            user.ID,
		Username:      user.Username,
		DisplayName:   user.DisplayName,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Permissions:   permissions,
	}
}
