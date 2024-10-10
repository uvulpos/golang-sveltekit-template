package models

import serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"

type UserModel struct {
	ID            string `db:"id"`
	Username      string `db:"username"`
	DisplayName   string `db:"display_name"`
	Email         string `db:"email"`
	EmailVerified bool   `db:"email_verified"`
}

func (m *UserModel) ToServiceModel() *serviceModel.UserModel {
	return &serviceModel.UserModel{
		ID:            m.ID,
		Username:      m.Username,
		DisplayName:   m.DisplayName,
		Email:         m.Email,
		EmailVerified: m.EmailVerified,
	}
}
