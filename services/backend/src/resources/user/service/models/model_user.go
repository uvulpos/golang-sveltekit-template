package models

type UserModel struct {
	ID            string
	Username      string
	DisplayName   string
	Email         string
	EmailVerified bool
}
