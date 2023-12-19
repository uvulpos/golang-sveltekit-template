package storage

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (h *UserStore) SetPasswordByID(uuid, newPassword string) error {

	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(newPassword), 8)
	if hashedPasswordErr != nil {
		return errors.New("could not update user")
	}

	rows := h.dbstore.DB.QueryRowx(
		"UPDATE users SET password=$2 WHERE id=$1 AND auth_source='basic'",
		uuid,
		hashedPassword,
	)
	if rows == nil {
		return errors.New("could not update user")
	}

	return nil
}
