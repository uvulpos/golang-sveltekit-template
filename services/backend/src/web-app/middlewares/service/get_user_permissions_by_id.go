package service

import "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"

func (s *MiddlewareService) GetUserPermissionsByID(userID string) ([]string, customerrors.ErrorInterface) {
	return s.userSvc.GetUserPermissionsByID(nil, userID)
}
