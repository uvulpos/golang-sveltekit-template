package storage

import (
	"errors"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

func (h *UserStore) GetUsers(tx *sqlx.Tx) ([]*service.User, error) {

	var rows *sqlx.Rows
	var rowErr error
	const sql = `SELECT * from public.user_with_rolename LIMIT 50`
	if tx == nil {
		rows, rowErr = h.dbstore.DB.Queryx(sql)
	} else {
		rows, rowErr = tx.Queryx(sql)
	}

	if rows == nil || rowErr != nil {
		return nil, errors.New("no user found")
	}

	var users []UserWithRole

	for rows.Next() {
		user := UserWithRole{}
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	svcUser := ToUserWithRoleSvcModels(users)
	return svcUser, nil
}
