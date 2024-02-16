package storage

import (
	"errors"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/resources/users/service"
)

func (h *UserStore) GetUserByUsername(tx *sqlx.Tx, username string) (*service.UserWithPermission, error) {
	var rows *sqlx.Rows
	var rowErr error
	if tx == nil {
		rows, rowErr = h.dbstore.DB.Queryx("SELECT * from public.user_with_permission WHERE username=$1 LIMIT 1", username)
	} else {
		rows, rowErr = tx.Queryx("SELECT * from public.user_with_permission WHERE username=$1 LIMIT 1", username)
	}
	if rowErr != nil {
		return nil, errors.New("no user found")
	}

	var users []UserWithPermission
	for rows.Next() {
		var user UserWithPermission
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, nil
	} else if len(users) > 1 {
		return nil, errors.New("too many users with same username found")
	}

	svcUser := users[0].ToUserWithPermissionsSvcModel()
	return svcUser, nil
}
