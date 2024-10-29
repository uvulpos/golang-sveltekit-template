package service

import (
	"context"
	"encoding/json"
	"io"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
	models "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/authentik/models"
	providerConst "github.com/uvulpos/golang-sveltekit-template/src/resources/auth/service/provider-const"
)

func (s *AuthService) AuthentikCallbackFunction(authCode, state string) (string, customerrors.ErrorInterface) {

	authToken, err := s.authentikConfig.Exchange(context.Background(), authCode)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "(oauth callback authentik) Failed to exchange auth code with authentik provider")
	}

	client := s.authentikConfig.Client(context.Background(), authToken)
	resp, err := client.Get(s.authentikOauthUserInfoEP)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "(oauth callback authentik) Failed to get user info from authentik provider")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "(oauth callback authentik) Failed to read authentik oauth user response body")
	}

	var result *models.AuthentikUserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return "", customerrors.NewInternalServerError(err, "", "(oauth callback authentik) Failed to unmarshal authentik oauth user response body")
	}

	loginUserID, loginUserError := s.userSvc.GetUserIDByLogin(providerConst.Authentik, result.ID)
	if loginUserError != nil {

		if loginUserError.ErrorType() != customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
			return "", loginUserError
		}

		// Create new user if user does not exist or relationship cannot be established
		tx, txErr := s.userSvc.StartTransaction()
		if txErr != nil {
			return "", txErr
		}

		defer func(tx *sqlx.Tx) {
			tx.Rollback()
		}(tx)

		createdUserID, createUserErr := s.userSvc.CreateUser(
			tx,
			result.Name,
			result.PreferredUsername,
			result.Email,
			result.EmailVerified,
		)
		if createUserErr != nil {
			tx.Rollback()
			return "", createUserErr
		}

		loginUserID = createdUserID

		createdUserLoginIdentityErr := s.userSvc.CreateUserLoginIdentity(
			tx,
			createdUserID,
			providerConst.Authentik,
			result.ID,
		)
		if createdUserLoginIdentityErr != nil {
			tx.Rollback()
			return "", createdUserLoginIdentityErr
		}

		commitErr := tx.Commit()
		if commitErr != nil {
			return "", customerrors.NewDatabaseError(commitErr, "", "(oauth callback authentik) Failed to commit transaction (create user)", "", nil)
		}
	}

	return loginUserID, nil
}
