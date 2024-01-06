package service

import (
	"strings"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/helper/check"
)

func (h *UserSvc) IsAvailableUsername(tx *sqlx.Tx, username, jwtUserID string) (bool, error) {

	var txStarted bool = false
	username = strings.ToLower(username)

	if tx == nil {
		tx, txErr := h.storage.CreateTransaction()
		if txErr != nil {
			return false, txErr
		}
		txStarted = true
		defer tx.Rollback()
	}

	isValid := check.CheckIsValidUsername(username)

	user, userErr := h.storage.GetUserByUsername(tx, username)
	if userErr != nil {
		return false, userErr
	}

	if txStarted && tx != nil {
		commitErr := tx.Commit()
		if commitErr != nil {
			return false, commitErr
		}
	}

	if isValid && user.IsSameUserOrUndefined(jwtUserID) {
		return true, nil
	}
	return false, nil
}
