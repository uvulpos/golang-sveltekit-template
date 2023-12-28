package service

import (
	"github.com/go-sqlx/sqlx"
	httpModels "github.com/uvulpos/go-svelte/src/resources/users/http/models"
)

type UserSvc struct {
	storage UserStorage
}

func NewUserSvc(storage UserStorage) *UserSvc {
	return &UserSvc{
		storage,
	}
}

type UserStorage interface {
	CreateTransaction() (*sqlx.Tx, error)
	GetUserByCredentials(username, password string) (*UserWithPermission, error)
	GetUserByUUID(uuid string) (*UserWithPermission, error)
	GetUserByUsername(tx *sqlx.Tx, username string) (*UserWithPermission, error)
	GetUserByEmail(tx *sqlx.Tx, email string) (*UserWithPermission, error)
	SetPasswordByID(uuid, newPassword string) error
	GetUserCountByCredentials(username string) (int, error)
	UpdateUserData(tx *sqlx.Tx, payload httpModels.ChangeUserDataPayload, userUuID string) error
}
