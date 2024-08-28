package service

import (
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *UserService) GetUserPermissionsByID(userID string) ([]string, customerrors.ErrorInterface) {
	return s.storage.GetUserPermissionsByID(nil, userID)
}
