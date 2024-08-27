package models

type UserModel struct {
	ID            string
	Username      string
	DisplayName   string
	Email         string
	EmailVerified bool
}

func NewUserModel(id, username, displayName, email string, emailVerified bool) *UserModel {
	return &UserModel{
		ID:            id,
		Username:      username,
		DisplayName:   displayName,
		Email:         email,
		EmailVerified: emailVerified,
	}
}
