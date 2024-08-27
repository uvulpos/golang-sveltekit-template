package models

import serviceModel "github.com/uvulpos/go-svelte/src/resources/user/service/models"

type UserModel struct {
	ID            string `db:"id"`
	Username      string `db:"username"`
	DisplayName   string `db:"display_name"`
	Email         string `db:"email"`
	EmailVerified bool   `db:"email_verified"`
}

func (m *UserModel) ToServiceModel() *serviceModel.UserModel {
	return serviceModel.NewUserModel(
		m.ID,
		m.Username,
		m.DisplayName,
		m.Email,
		m.EmailVerified,
	)
}
