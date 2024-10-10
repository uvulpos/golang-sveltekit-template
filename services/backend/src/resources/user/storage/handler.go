package storage

import (
	"encoding/json"
	"errors"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db,
	}
}

func (s *UserStore) StartTransaction() (*sqlx.Tx, customerrors.ErrorInterface) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, customerrors.NewDatabaseTransactionBeginError(err, "Failed to start transaction")
	}
	return tx, nil
}

func (permission *Permission) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &permission)
}

func (s *UserPermissions) ToServiceModel() *serviceModel.UserPermissionsModel {
	return &serviceModel.UserPermissionsModel{
		Permissions: PermissionsToStringArray(s.Permissions),
	}
}

func PermissionsToStringArray(input []Permission) []string {
	var permissionArray []string

	for _, permission := range input {
		permissionArray = append(permissionArray, string(permission))
	}

	return permissionArray
}
