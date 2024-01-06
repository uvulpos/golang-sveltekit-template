package storage

import (
	"errors"

	"github.com/go-sqlx/sqlx"
	httpModels "github.com/uvulpos/go-svelte/src/resources/users/http/models"
)

func (h *UserStore) UpdateUserData(tx *sqlx.Tx, payload httpModels.ChangeUserDataPayload, userUuid string) error {
	var rowErr error
	const sql = `UPDATE users SET username = $1, email=$2 WHERE id=$3`
	if tx == nil {
		_, rowErr = h.dbstore.DB.Exec(sql,
			payload.Username,
			payload.Email,
			userUuid,
		)
	} else {
		_, rowErr = tx.Exec(sql,
			payload.Username,
			payload.Email,
			userUuid,
		)
	}

	if rowErr != nil {
		return errors.New("no user found")
	}

	return nil
}
