package http

import (
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

type UserWithPermission struct {
	Id                  string          `json:"uuid"`
	Username            string          `json:"username"`
	Email               string          `json:"email"`
	AuthSource          string          `json:"auth_source"`
	AdminReviewRequired bool            `json:"suspendet"`
	RoleName            string          `json:"role_name"`
	Permissions         UserPermissions `json:"permissions"`
}

type UserPermissions []UserPermission
type UserPermission struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

func ToUserWithPermission(serviceUser *service.UserWithPermission) *UserWithPermission {
	permissions := toPermissions(serviceUser.Permissions)

	return &UserWithPermission{
		Id:                  serviceUser.Id.String(),
		Username:            serviceUser.Username,
		Email:               serviceUser.Email,
		AuthSource:          serviceUser.AuthSource,
		AdminReviewRequired: serviceUser.AdminReviewRequired,
		RoleName:            serviceUser.RoleName,
		Permissions:         permissions,
	}
}

func toPermissions(servicePermissions service.UserPermissions) []UserPermission {
	var permissions []UserPermission
	for _, sPermission := range servicePermissions {
		permissions = append(permissions, UserPermission{
			Name:       sPermission.Name,
			Identifier: sPermission.Identifier,
		})
	}

	return permissions
}
