package storage

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

type UserWithPermissions []UserWithPermission
type UserWithPermission struct {
	Id                  uuid.UUID `db:"id"`
	Username            string    `db:"username"`
	Email               string    `db:"email"`
	LdapUUID            string    `db:"ldap_uuid"`
	AuthSource          string    `db:"auth_source"`
	AdminReviewRequired bool      `db:"admin_review_required"`
	RoleID              string    `db:"role_id"`
	RoleName            string    `db:"role_name"`
	Permissions         string    `db:"permissions"`
}

func ToUserWithPermissionsSvcModels(u UserWithPermissions) []*service.UserWithPermission {
	var users []*service.UserWithPermission
	for _, user := range u {
		users = append(users, user.ToUserWithPermissionsSvcModel())
	}
	return users
}

func (u UserWithPermission) ToUserWithPermissionsSvcModel() *service.UserWithPermission {
	return &service.UserWithPermission{
		Id:                  u.Id,
		Username:            u.Username,
		Email:               u.Email,
		LdapUUID:            u.LdapUUID,
		AuthSource:          u.AuthSource,
		AdminReviewRequired: u.AdminReviewRequired,
		RoleID:              u.RoleID,
		RoleName:            u.RoleName,
		Permissions:         u.Permissions,
	}
}
