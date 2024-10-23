package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (s *UserService) CreateUserLoginIdentity(
	tx *sqlx.Tx,
	createdUserID string,
	providerType providerConst.AuthProvider,
	providerID string,
) customerrors.ErrorInterface {
	return s.storage.CreateUserLoginIdentity(tx, createdUserID, string(providerType), providerID)
}
