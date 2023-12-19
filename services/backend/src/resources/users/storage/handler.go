package storage

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/uvulpos/go-svelte/src/helper/database"
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

type UserStore struct {
	dbstore database.Sql
}

func NewUserStore(db database.Sql) *UserStore {
	return &UserStore{
		db,
	}
}

type UserWithPermissions []UserWithPermission
type UserWithPermission struct {
	Id                  uuid.UUID       `db:"id"`
	Username            string          `db:"username"`
	Email               string          `db:"email"`
	Password            sql.NullString  `db:"password"`
	LdapUUID            sql.NullString  `db:"ldap_uuid"`
	AuthSource          string          `db:"auth_source"`
	AdminReviewRequired bool            `db:"admin_review_required"`
	RoleID              string          `db:"role_id"`
	RoleName            string          `db:"role_name"`
	Permissions         UserPermissions `db:"permissions"`
}

type UserPermissions []UserPermission
type UserPermission struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Identifier  string    `json:"identifier" db:"identifier"`
}

func ToUserWithPermissionsSvcModels(u UserWithPermissions) []*service.UserWithPermission {
	var users []*service.UserWithPermission
	for _, user := range u {
		users = append(users, user.ToUserWithPermissionsSvcModel())
	}
	return users
}

func (u UserWithPermission) ToUserWithPermissionsSvcModel() *service.UserWithPermission {
	var ldapUuid *string = nil
	if u.LdapUUID.Valid {
		ldapUuid = &u.LdapUUID.String
	}
	return &service.UserWithPermission{
		Id:                  u.Id,
		Username:            u.Username,
		Email:               u.Email,
		LdapUUID:            ldapUuid,
		AuthSource:          u.AuthSource,
		AdminReviewRequired: u.AdminReviewRequired,
		RoleID:              u.RoleID,
		RoleName:            u.RoleName,
		Permissions:         toSvcPermissions(u.Permissions),
	}
}

func toSvcPermissions(perms UserPermissions) service.UserPermissions {
	var permissions service.UserPermissions

	for _, p := range perms {
		permissions = append(permissions, service.UserPermission{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Identifier:  p.Identifier,
		})
	}

	return permissions
}

func (up *UserPermissions) Value() (driver.Value, error) {
	return json.Marshal(up)
}

func (up *UserPermissions) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, up)
	case string:
		return json.Unmarshal([]byte(v), up)
	}
	return errors.New("failed to assert db type")
}
