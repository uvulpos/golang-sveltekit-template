package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"

	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
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
	CreateUser(tx *sqlx.Tx, displayName string, username string, email string, emailVerified bool) (string, customerrors.ErrorInterface)
	CreateUserLoginIdentity(tx *sqlx.Tx, createdUserID string, authProvider string, authProviderID string) customerrors.ErrorInterface
	GetUserPermissionsByID(tx *sqlx.Tx, userID string) ([]string, customerrors.ErrorInterface)
	GetUserByID(tx *sqlx.Tx, userID string) (*serviceModel.UserModel, customerrors.ErrorInterface)
	GetUserByUsername(tx *sqlx.Tx, username string) (*serviceModel.UserModel, customerrors.ErrorInterface)
	GetUserAuthSessionByID(tx *sqlx.Tx, sessionID string) (*serviceModel.SessionModel, customerrors.ErrorInterface)
	GetUserIDByLogin(provider string, providerID string) (string, customerrors.ErrorInterface)
}

func (s *UserService) StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface) {
	return s.storage.StartTransaction()
}
