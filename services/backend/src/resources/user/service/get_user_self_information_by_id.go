package service

import (
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

func (s *UserService) GetUserSelfInformationByID(userID string) (*serviceModel.UserSelfInformationModel, customerrors.ErrorInterface) {

	// get normal user entity
	user, userErr := s.storage.GetUserByID(nil, userID)
	if userErr != nil {
		return nil, userErr
	}

	permissions, permissionsErr := s.GetUserPermissionsByID(nil, userID)
	if permissionsErr != nil {
		return nil, permissionsErr
	}

	return serviceModel.NewUserSelfInformationModel(user, permissions), nil
}
