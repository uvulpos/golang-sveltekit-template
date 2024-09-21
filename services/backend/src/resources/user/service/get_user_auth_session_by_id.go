package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"

	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
)

func (s *UserService) GetUserAuthSessionByID(tx *sqlx.Tx, sessionID string) (*serviceModel.SessionModel, customerrors.ErrorInterface) {
	return s.storage.GetUserAuthSessionByID(tx, sessionID)
}
