package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	customerrorconst "github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors/custom-error-const"
)

func (s *AuthService) AuthentikCallbackFunction(authCode, state, oauthUserinfoURL string) (string, customerrors.ErrorInterface) {

	authToken, err := s.authentikConfig.Exchange(context.Background(), authCode)
	if err != nil {
		fmt.Println("Exchange ERROR")
		return "", customerrors.NewInternalServerError(err, "", "Failed to exchange auth code with authentik provider")
	}

	client := s.authentikConfig.Client(context.Background(), authToken)
	resp, err := client.Get(s.authentikOauthUserInfoEP)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "Failed to get user info from authentik provider")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "Failed to read authentik oauth user response body")
	}

	fmt.Println("Data: ", string(data))

	var result *AuthentikUserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return "", customerrors.NewInternalServerError(err, "", "Failed to unmarshal authentik oauth user response body")
	}

	loginUserID, loginUserError := s.storage.GetUserIDByLogin("Authentik", result.ID)
	if loginUserError != nil {

		if loginUserError.ErrorType() != customerrorconst.ERROR_IDENTIFIER_DATABASE_NOT_FOUND {
			return "", loginUserError
		}

		// Create new user if user does not exist or relationship cannot be established
		tx, txErr := s.storage.StartTransaction()
		if txErr != nil {
			return "", txErr
		}

		defer func(tx *sqlx.Tx) {
			tx.Rollback()
		}(tx)

		createdUserID, createUserErr := s.storage.CreateUser(
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

		createdUserLoginIdentityErr := s.storage.CreateUserLoginIdentity(
			tx,
			createdUserID,
			"Authentik",
			result.ID,
		)
		if createdUserLoginIdentityErr != nil {
			tx.Rollback()
			return "", createdUserLoginIdentityErr
		}

		commitErr := tx.Commit()
		if commitErr != nil {
			return "", customerrors.NewDatabaseError(err, "", "Failed to commit transaction", "", nil)
		}
	}

	return loginUserID, nil
}
