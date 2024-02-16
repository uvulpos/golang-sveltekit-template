package storage

import (
	"errors"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

func (h *UserStore) GetUserByUUID(tx *sqlx.Tx, uuid string) (*service.UserWithPermission, error) {

	var rows *sqlx.Row
	const sql = `SELECT * from public.user_with_permission WHERE id=$1 LIMIT 1`
	if tx == nil {
		rows = h.dbstore.DB.QueryRowx(sql, uuid)
	} else {
		rows = tx.QueryRowx(sql, uuid)
	}

	if rows == nil {
		return nil, errors.New("no user found")
	}

	user := &UserWithPermission{}
	scanErr := rows.StructScan(user)
	if scanErr != nil {
		return nil, scanErr
	}

	svcUser := user.ToUserWithPermissionsSvcModel()
	return svcUser, nil
}
