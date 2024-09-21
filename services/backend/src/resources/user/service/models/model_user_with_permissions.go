package models

type UserWithPermissionsModel struct {
	ID            string
	Username      string
	DisplayName   string
	Email         string
	EmailVerified bool
	Permissions   []string
}
