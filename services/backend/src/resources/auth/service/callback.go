package service

import (
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	badrequestconstraints "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/bad-request-constraints"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
	jwtService "github.com/uvulpos/golang-sveltekit-template/src/resources/jwt/service"
)

func (s *AuthService) CallbackFunction(provider, authCode, state string) (string, string, customerrors.ErrorInterface) {
	var loggedinUser string
	var loggedinUserErr customerrors.ErrorInterface

	if provider == string(providerConst.Authentik) && configuration.AUTHORIZATION_OAUTH_AUTHENIK {
		loggedinUser, loggedinUserErr = s.authentikProviderSvc.AuthentikCallbackFunction(authCode, state)
		if loggedinUserErr != nil {
			return "", "", loggedinUserErr
		}
	} else {
		return "", "", customerrors.NewBadRequestError(badrequestconstraints.BadRequest_AuthenticationProviderInvalid)
	}

	tx, txErr := s.storage.StartTransaction()
	if txErr != nil {
		return "", "", txErr
	}

	defer tx.Rollback()

	// as long as they are not used, they are nil / null by default
	var dbIPAddr *string = nil
	var dbUserAgent *string = nil

	sessionID, sessionIDErr := s.storage.StartLoginSession(tx, loggedinUser, dbUserAgent, dbIPAddr)
	if sessionIDErr != nil {
		return "", "", sessionIDErr
	}

	permissionScopes, permissionScopesErr := s.userSvc.GetUserPermissionsByID(tx, loggedinUser)
	if permissionScopesErr != nil {
		return "", "", permissionScopesErr
	}

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
