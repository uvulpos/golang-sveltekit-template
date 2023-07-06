package storage

import (
	"database/sql"

	"github.com/google/uuid"
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
	Id                    uuid.UUID `db:"id"`
	username              string    `db:"username"`
	email                 string    `db:"email"`
	ldap_uuid             string    `db:"ldap_uuid"`
	auth_source           string    `db:"auth_source"`
	admin_review_required bool      `db:"admin_review_required"`
	role_id               string    `db:"role_id"`
	role_name             string    `db:"role_name"`
	permissions           string    `db:"permissions"`
}
