package service

import (
	"fmt"

	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	"github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
)

func (s *AuthService) RecreateJwtFromSession(sessionID string) (string, customerrors.ErrorInterface) {
	session, sessionErr := s.userSvc.GetUserAuthSessionByID(nil, sessionID)
	if sessionErr != nil {
		if sessionErr.ErrorType() == customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
			return "", customerrors.NewNotAuthorizedError()
		}
	}

	user, userErr := s.userSvc.GetUserByID(nil, session.UserID)
	if userErr != nil {
		return "", userErr
	}

	permissions, permissionsErr := s.userSvc.GetUserPermissionsByID(nil, user.ID)
	if permissionsErr != nil {
		return "", permissionsErr
	}

	fmt.Println("CREATE JWT FROM SESSION")
	fmt.Println("CREATE JWT FROM SESSION")
	fmt.Println("CREATE JWT FROM SESSION")
	fmt.Println("CREATE JWT FROM SESSION")
	fmt.Println("permissions", permissions)

	refreshToken, refreshTokenErr := s.jwt.CreateJWT(service.NewJwtDataModel(user.ID, sessionID, permissions))
	if refreshTokenErr != nil {
		return "", refreshTokenErr
	}

	return refreshToken, nil
}
