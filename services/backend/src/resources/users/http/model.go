package http

import (
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

type AdvancedUser struct {
	Id                  string `json:"uuid"`
	Username            string `json:"username"`
	Email               string `json:"email"`
	AuthSource          string `json:"auth_source"`
	AdminReviewRequired bool   `json:"suspendet"`
	RoleName            string `json:"role_name"`
}

type AdvancedUserWithPermission struct {
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

func ToUserForManager(serviceUser *service.User) *AdvancedUser {
	return &AdvancedUser{
		Id:                  serviceUser.Id.String(),
		Username:            serviceUser.Username,
		Email:               serviceUser.Email,
		AuthSource:          serviceUser.AuthSource,
		AdminReviewRequired: serviceUser.AdminReviewRequired,
		RoleName:            serviceUser.Role.Name,
	}

}

func ToUsersForManager(serviceUsers []*service.User) []*AdvancedUser {
	var users []*AdvancedUser
	for _, serviceUser := range serviceUsers {
		users = append(users, ToUserForManager(serviceUser))
	}
	return users
}

func ToUserWithPermissionsForManager(serviceUser *service.UserWithPermission) *AdvancedUserWithPermission {

	var permissions []UserPermission
	for _, permission := range serviceUser.Permissions {
		permissions = append(permissions, UserPermission{
			Name:       permission.Name,
			Identifier: permission.Identifier,
		})
	}

	return &AdvancedUserWithPermission{
		Id:                  serviceUser.Id.String(),
		Username:            serviceUser.Username,
		Email:               serviceUser.Email,
		AuthSource:          serviceUser.AuthSource,
		AdminReviewRequired: serviceUser.AdminReviewRequired,
		RoleName:            serviceUser.Role.Name,
		Permissions:         permissions,
	}

}

func ToUsersWithPermissionsForManager(serviceUsers []*service.UserWithPermission) []*AdvancedUserWithPermission {
	var users []*AdvancedUserWithPermission
	for _, serviceUser := range serviceUsers {
		users = append(users, ToUserWithPermissionsForManager(serviceUser))
	}
	return users
}

func ToSelfUserWithPermission(serviceUser *service.UserWithPermission) *AdvancedUserWithPermission {
	permissions := toPermissions(serviceUser.Permissions)

	return &AdvancedUserWithPermission{
		Id:                  serviceUser.Id.String(),
		Username:            serviceUser.Username,
		Email:               serviceUser.Email,
		AuthSource:          serviceUser.AuthSource,
		AdminReviewRequired: serviceUser.AdminReviewRequired,
		RoleName:            serviceUser.Role.Name,
		Permissions:         permissions,
	}
}

func ToUserWithPermission(serviceUser *service.UserWithPermission) *AdvancedUserWithPermission {
	permissions := toPermissions(serviceUser.Permissions)

	return &AdvancedUserWithPermission{
		Id:                  serviceUser.Id.String(),
		Username:            serviceUser.Username,
		Email:               serviceUser.Email,
		AuthSource:          serviceUser.AuthSource,
		AdminReviewRequired: serviceUser.AdminReviewRequired,
		RoleName:            serviceUser.Role.Name,
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
