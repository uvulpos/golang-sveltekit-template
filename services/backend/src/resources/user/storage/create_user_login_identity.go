package storage

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *UserStore) CreateUserLoginIdentity(tx *sqlx.Tx, createdUserID string, authProvider string, authProviderID string) customerrors.ErrorInterface {
	sqlquery := "INSERT INTO public.user_identities (provider, provider_user_id, user_id) VALUES ($1, $2, $3)"

	_, err := tx.Exec(sqlquery, authProvider, authProviderID, createdUserID)
	if err != nil {
		data := customerrors.SqlData{
			"authProvider":   authProvider,
			"authProviderID": authProviderID,
			"userID":         createdUserID,
		}

		return customerrors.NewDatabaseError(err, "", "", sqlquery, data)
	}
	return nil
}
