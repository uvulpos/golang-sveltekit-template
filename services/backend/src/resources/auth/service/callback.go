package service

import (
	"fmt"

	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
)

func (s *AuthService) CallbackFunction(authCode, state, ipaddr, userAgent string) (string, string, customerrors.ErrorInterface) {
	loggedinUser, loggedinUserErr := s.auth.AuthentikCallbackFunction(authCode, state)
	if loggedinUserErr != nil {
		return "", "", loggedinUserErr
	}

	tx, txErr := s.storage.StartTransaction()
	if txErr != nil {
		return "", "", txErr
	}

	defer tx.Rollback()

	var dbIPAddr *string = nil
	if ipaddr != "" {
		dbIPAddr = &ipaddr
	}

	var dbUserAgent *string = nil
	if userAgent != "" {
		dbUserAgent = &userAgent
	}

	sessionID, sessionIDErr := s.storage.StartLoginSession(tx, loggedinUser, dbUserAgent, dbIPAddr)
	if sessionIDErr != nil {
		return "", "", sessionIDErr
	}

	permissionScopes, permissionScopesErr := s.userSvc.GetUserPermissionsByID(tx, loggedinUser)
	if permissionScopesErr != nil {
		return "", "", permissionScopesErr
	}

	fmt.Println("Permissions")
	fmt.Println("Permissions")
	fmt.Println("Permissions")
	fmt.Println("Permissions")
	fmt.Println(permissionScopes)

	commitErr := tx.Commit()
	if commitErr != nil {
		return "", "", customerrors.NewDatabaseTransactionCommitError(commitErr, "Failed to commit transaction")
	}

	jwt, jwtErr := s.jwt.CreateJWT(jwtService.NewJwtDataModel(loggedinUser, sessionID, permissionScopes))
	if jwtErr != nil {
		return "", "", customerrors.NewInternalServerError(jwtErr, loggedinUser, "cannot create jwt after login / signup")
	}

	refreshToken, refreshTokenErr := s.jwt.CreateRefreshToken(sessionID)
	if refreshTokenErr != nil {
		return "", "", customerrors.NewInternalServerError(refreshTokenErr, loggedinUser, "cannot create refreshToken after login / signup")
	}

	return jwt, refreshToken, nil
}
