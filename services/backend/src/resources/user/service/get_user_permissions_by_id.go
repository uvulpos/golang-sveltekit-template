package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *UserService) GetUserPermissionsByID(tx *sqlx.Tx, userID string) ([]string, customerrors.ErrorInterface) {
	return s.storage.GetUserPermissionsByID(tx, userID)
}
