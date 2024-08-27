package httpModels

import (
	serviceModel "github.com/uvulpos/go-svelte/src/resources/user/service/models"
)

type SelfInformationModel struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	DisplayName string   `json:"display_name"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
}

func NewSelfInformationModel(m *serviceModel.UserSelfInformationModel) *SelfInformationModel {
	return &SelfInformationModel{
		ID:          m.User.ID,
		Username:    m.User.Username,
		DisplayName: m.User.DisplayName,
		Email:       m.User.Email,
		Permissions: m.Permissions,
	}
}
