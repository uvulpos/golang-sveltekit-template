package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

func (s *UserService) GetUserByUsernameWithPermissions(tx *sqlx.Tx, username string) (*models.UserWithPermissionsModel, customerrors.ErrorInterface) {
	user, userErr := s.GetUserByUsername(tx, username)
	if userErr != nil {
		return nil, userErr
	}

	permissions, permissionsErr := s.GetUserPermissionsByID(tx, user.ID)
	if permissionsErr != nil {
		return nil, permissionsErr
	}

	return models.NewUserWithPermissionsModel(user, permissions), nil
}
