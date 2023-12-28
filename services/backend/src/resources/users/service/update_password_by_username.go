package service

import (
	"errors"
)

func (h *UserSvc) UpdatePasswordByUsername(username, newPassword, oldPassword string) error {

	if newPassword == "" || oldPassword == "" {
		return errors.New("password cannot be empty")
	}

	user, userAuth := h.GetAuthenticatedUserByCredentials(username, oldPassword)
	if userAuth != nil {
		return userAuth
	}

	if user.AuthSource != "basic" {
		return errors.New("only basic authenticated users can change their password")
	}

	passwordChangeErr := h.storage.SetPasswordByID(user.Id.String(), newPassword)
	if passwordChangeErr != nil {
		return passwordChangeErr
	}

	return nil
}
