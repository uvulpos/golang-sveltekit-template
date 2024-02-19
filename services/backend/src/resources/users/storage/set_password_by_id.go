package storage

import (
	"errors"

	"github.com/go-sqlx/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserStore) SetPasswordByID(tx *sqlx.Tx, uuid, newPassword string) error {

	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(newPassword), 8)
	if hashedPasswordErr != nil {
		return errors.New("could not update user")
	}

	sql := `UPDATE users SET password=$2 WHERE id=$1 AND auth_source='basic'`

	var rows *sqlx.Row
	if tx == nil {
		rows = h.dbstore.DB.QueryRowx(sql, uuid, hashedPassword)
	} else {
		rows = tx.QueryRowx(sql, uuid, hashedPassword)
	}
	if rows == nil {
		return errors.New("could not update user")
	}

	return nil
}

// UPDATE users SET password='$2a$08$o1vsZKCXhOPXoNBprne0ceI.qEnK9tMitP9tcf3Bk/M7AvgjrNPmS' WHERE id='4c7537f0-45b7-44da-8955-fd3d28ea58d8'
