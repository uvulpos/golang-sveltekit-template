package service

import (
	"errors"

	"github.com/go-sqlx/sqlx"
)

func (h *UserSvc) UpdatePassword(tx *sqlx.Tx, username, newPassword, oldPassword string) error {
	if newPassword == "" || oldPassword == "" {
		return errors.New("password cannot be empty")
	}

	var txStarted bool = false
	if tx == nil {
		tx, txErr := h.storage.CreateTransaction()
		if txErr != nil {
			return txErr
		}
		txStarted = true
		defer tx.Rollback()
	}

	tx, txErr := h.storage.CreateTransaction()
	if txErr != nil {
		return txErr
	}
	defer tx.Rollback()

	user, userErr := h.GetAuthenticatedUserByCredentials(tx, username, oldPassword)
	if userErr != nil {
		return userErr
	}

	if user.AuthSource != "basic" {
		return errors.New("only basic authenticated users can change their password")
	}

	passwordChangeErr := h.storage.SetPasswordByID(tx, user.Id.String(), newPassword)
	if passwordChangeErr != nil {
		return passwordChangeErr
	}

	if txStarted && tx != nil {
		commitErr := tx.Commit()
		if commitErr != nil {
			return commitErr
		}
	}

	return nil
}
