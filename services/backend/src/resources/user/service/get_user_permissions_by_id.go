package service

import (
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

func (s *UserService) GetUserPermissionsByID(userID string) ([]string, customerrors.ErrorInterface) {
	return s.storage.GetUserPermissionsByID(nil, userID)
}
