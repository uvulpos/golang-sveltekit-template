package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

type UserService struct {
	storage AuthStorageInterface
}

func NewUserService(storage AuthStorageInterface) *UserService {
	return &UserService{
		storage: storage,
	}
}

type AuthStorageInterface interface {
	StartTransaction() (tx *sqlx.Tx, err customerrors.ErrorInterface)
	GetUserPermissionsByID(tx *sqlx.Tx, userID string) ([]string, customerrors.ErrorInterface)
}
