package service

import "github.com/go-sqlx/sqlx"

func (h *UserSvc) GetUserByUUID(tx *sqlx.Tx, uuid string) (*UserWithPermission, error) {
	return h.storage.GetUserByUUID(tx, uuid)

}
