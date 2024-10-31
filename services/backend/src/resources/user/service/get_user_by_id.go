package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

func (s *UserService) GetUserByID(tx *sqlx.Tx, userID string) (*models.UserModel, customerrors.ErrorInterface) {
	user, userErr := s.storage.GetUserByID(tx, userID)
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}
