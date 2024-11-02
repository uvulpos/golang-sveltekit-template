package storage

import (
	"database/sql"

	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (s *UserStore) GetUserIDByLogin(provider string, providerID string) (string, customerrors.ErrorInterface) {
	var userID string
	const sqlquery = "SELECT user_id FROM public.user_identities WHERE provider=$1 AND provider_user_id=$2"

	err := s.db.Get(&userID, sqlquery, provider, providerID)
	if err != nil {

		data := customerrors.SqlData{
			"provider":         provider,
			"provider_user_id": providerID,
		}

		switch err {

		case sql.ErrNoRows:
			return "", customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return "", customerrors.NewDatabaseError(err, "", "", sqlquery, nil)
		}
	}

	return userID, nil
}
