package service

import (
	"github.com/go-sqlx/sqlx"
	jwtService "github.com/uvulpos/go-svelte/authentication-api/ressources/jwt/service"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

func (s *AuthService) CallbackFunction(authCode, state, oauthUserinfoURL, ipaddr, userAgent string) (string, customerrors.ErrorInterface) {
	loggedinUser, loggedinUserErr := s.auth.AuthentikCallbackFunction(authCode, state, oauthUserinfoURL)
	if loggedinUserErr != nil {
		return "", loggedinUserErr
	}

	tx, txErr := s.storage.StartTransaction()
	if txErr != nil {
		return "", txErr
	}

	defer func(tx *sqlx.Tx) {
		tx.Rollback()
	}(tx)

	var dbIPAddr *string = nil
	if ipaddr != "" {
		dbIPAddr = &ipaddr
	}

	var dbUserAgent *string = nil
	if userAgent != "" {
		dbUserAgent = &userAgent
	}

	sessionID, sessinErr := s.storage.StartLoginSession(tx, loggedinUser, dbUserAgent, dbIPAddr)
	if sessinErr != nil {
		return "", sessinErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return "", customerrors.NewDatabaseTransactionCommitError(commitErr, "Failed to commit transaction")
	}

	jwt, jwtErr := s.jwt.CreateJWT(jwtService.NewJwtDataModel(loggedinUser, sessionID))
	if jwtErr != nil {
		return "", customerrors.NewInternalServerError(jwtErr, loggedinUser, "cannot create jwt after login / signup")
	}

	return jwt, nil
}
