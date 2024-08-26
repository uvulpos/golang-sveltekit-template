package service

import "github.com/uvulpos/go-svelte/basic-utils/customerrors"

func (s *UserService) GetUserSelfInformationByID(userID string) (interface{}, customerrors.ErrorInterface) {

	_, permissionsErr := s.GetUserPermissionsByID(userID)
	if permissionsErr != nil {
		return nil, permissionsErr
	}

	return nil, nil
}
