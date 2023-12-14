package service

import (
	"github.com/google/uuid"
)

type UserWithPermission struct {
	Id                  uuid.UUID
	Username            string
	Email               string
	LdapUUID            *string
	AuthSource          string
	AdminReviewRequired bool
	RoleID              string
	RoleName            string
	Permissions         UserPermissions
}

type UserPermissions []UserPermission
type UserPermission struct {
	ID          uuid.UUID
	Name        string
	Description string
	Identifier  string
}
