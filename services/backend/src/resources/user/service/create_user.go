package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *UserService) CreateUser(tx *sqlx.Tx, name, prefName, email string, emailVerified bool) (string, customerrors.ErrorInterface) {
	return "ID", nil
}
