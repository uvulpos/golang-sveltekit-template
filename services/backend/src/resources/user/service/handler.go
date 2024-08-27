package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"

	serviceModel "github.com/uvulpos/go-svelte/src/resources/user/service/models"
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
	GetUserByID(tx *sqlx.Tx, userID string) (*serviceModel.UserModel, customerrors.ErrorInterface)
}
