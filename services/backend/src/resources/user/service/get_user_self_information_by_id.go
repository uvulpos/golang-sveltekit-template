package service

import (
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
	serviceModel "github.com/uvulpos/go-svelte/src/resources/user/service/models"
)

func (s *UserService) GetUserSelfInformationByID(userID string) (*serviceModel.UserSelfInformationModel, customerrors.ErrorInterface) {

	// get normal user entity
	user, userErr := s.storage.GetUserByID(nil, userID)
	if userErr != nil {
		return nil, userErr
	}

	permissions, permissionsErr := s.GetUserPermissionsByID(userID)
	if permissionsErr != nil {
		return nil, permissionsErr
	}

	selfInformationModel := serviceModel.NewUserSelfInformationModel(user, permissions)

	return selfInformationModel, nil
}