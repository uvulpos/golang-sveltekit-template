package service

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	jwt "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt"
	jwtModels "github.com/uvulpos/golang-sveltekit-template/src/helper/jwt/models"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (s *AuthService) RecreateJwtFromSession(sessionID string) (string, customerrors.ErrorInterface) {

	tx, txErr := s.storage.StartTransaction()
	if txErr != nil {
		return "", txErr
	}
	defer func(tx *sqlx.Tx) {
		tx.Rollback()
	}(tx)

	session, sessionErr := s.userSvc.GetUserAuthSessionByID(tx, sessionID)
	if sessionErr != nil {
		if sessionErr.ErrorType() == customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
			return "", customerrors.NewNotAuthorizedError()
		}
	}

	user, userErr := s.userSvc.GetUserByID(tx, session.UserID)
	if userErr != nil {
		return "", userErr
	}

	permissions, permissionsErr := s.userSvc.GetUserPermissionsByID(tx, user.ID)
	if permissionsErr != nil {
		return "", permissionsErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return "", customerrors.NewDatabaseError(commitErr, "", "(session jwt refresh) Failed to commit transaction (receive session + user data)", "", nil)
	}

	refreshToken, refreshTokenErr := jwt.CreateJWT(provider.Authentik, jwtModels.NewJwtDataModel(user.ID, sessionID, permissions))
	if refreshTokenErr != nil {
		return "", refreshTokenErr
	}

	return refreshToken, nil
}
