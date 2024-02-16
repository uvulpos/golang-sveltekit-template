package service

import (
	"github.com/google/uuid"
)

type User struct {
	Id                  uuid.UUID
	Username            string
	Email               string
	LdapUUID            *string
	AuthSource          string
	AdminReviewRequired bool
	Role                *UserRole
}

type UserRole struct {
	Id   string
	Name string
}

type UserWithPermission struct {
	Id                  uuid.UUID
	Username            string
	Email               string
	LdapUUID            *string
	AuthSource          string
	AdminReviewRequired bool
	Role                *UserRole
	Permissions         UserPermissions
}

type UserPermissions []UserPermission
type UserPermission struct {
	ID          uuid.UUID
	Name        string
	Description string
	Identifier  string
}

func (u UserWithPermission) IsSameUser(uuid string) bool {
	return u.Id.String() == uuid
}

func (u *UserWithPermission) IsSameUserOrUndefined(uuid string) bool {
	if u == nil {
		return true
	}
	return u.Id.String() == uuid
}

func (up UserPermissions) IncludePermission(permissionIdentifier string) bool {
	for _, permission := range up {
		if permission.Identifier == permissionIdentifier {
			return true
		}
	}
	return false
}
