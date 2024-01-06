package service

import (
	"strings"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/helper/check"
)

func (h *UserSvc) IsAvailableEmail(tx *sqlx.Tx, emailaddr, jwtUserID string) (bool, error) {

	var txStarted bool = false
	emailaddr = strings.ToLower(emailaddr)

	if tx == nil {
		tx, txErr := h.storage.CreateTransaction()
		if txErr != nil {
			return false, txErr
		}
		txStarted = true
		defer tx.Rollback()
	}

	isValid, isValidErr := check.CheckIsValidEmail(emailaddr)
	if isValidErr != nil {
		return false, isValidErr
	}

	user, userErr := h.storage.GetUserByEmail(tx, emailaddr)
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
