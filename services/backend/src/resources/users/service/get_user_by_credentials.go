package service

import "github.com/go-sqlx/sqlx"

func (h *UserSvc) GetAuthenticatedUserByCredentials(tx *sqlx.Tx, username, password string) (*UserWithPermission, error) {
	return h.storage.GetUserByCredentials(tx, username, password)
}
