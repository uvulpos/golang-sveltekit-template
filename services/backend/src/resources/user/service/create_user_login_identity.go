package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (*UserService) CreateUserLoginIdentity(
	tx *sqlx.Tx,
	createdUserID string,
	providerType providerConst.AuthProvider,
	providerID string,
) customerrors.ErrorInterface {
	return nil
}
