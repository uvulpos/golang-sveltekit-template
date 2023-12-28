package service

import (
	"errors"

	httpModels "github.com/uvulpos/go-svelte/src/resources/users/http/models"
)

func (h *UserSvc) UpdateUserData(payload httpModels.ChangeUserDataPayload, userUuid string) (*UserWithPermission, error) {
	tx, txErr := h.storage.CreateTransaction()
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback()

	usernameAvailable, usernameAvailableErr := h.IsAvailableUsername(tx, payload.Username, userUuid)
	if usernameAvailableErr != nil {
		return nil, usernameAvailableErr
	} else if !usernameAvailable {
		return nil, errors.New("username is invalid")
	}

	emailAvailable, emailAvailableErr := h.IsAvailableEmail(tx, payload.Email, userUuid)
	if emailAvailableErr != nil {
		return nil, emailAvailableErr
	} else if !emailAvailable {
		return nil, errors.New("email is invalid")
	}

	updatedUserErr := h.storage.UpdateUserData(tx, payload, userUuid)
	if updatedUserErr != nil {
		return nil, updatedUserErr
	}

	updatedUser, updatedUserErr := h.storage.GetUserByUUID(userUuid)
	if updatedUserErr != nil {
		return nil, updatedUserErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return nil, commitErr
	}

	return updatedUser, nil
}
