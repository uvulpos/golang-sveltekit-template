package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *UserService) CreateUser(tx *sqlx.Tx, name, username, email string, emailVerified bool) (string, customerrors.ErrorInterface) {
	return s.storage.CreateUser(tx, name, username, email, emailVerified)
}
