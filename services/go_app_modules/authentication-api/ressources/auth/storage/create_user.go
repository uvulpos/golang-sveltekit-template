package storage

import (
	"database/sql"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

func (s *AuthStorage) CreateUser(tx *sqlx.Tx, displayName string, username string, email string, emailVerified bool) (string, customerrors.ErrorInterface) {
	sqlquery := "INSERT INTO public.users (display_name, username, email, email_verified) VALUES ($1, $2, $3, $4) RETURNING id"

	var userID string
	err := tx.Get(&userID, sqlquery, displayName, username, email, emailVerified)
	if err != nil {

		data := customerrors.SqlData{
			"displayName":   displayName,
			"username":      username,
			"email":         email,
			"emailVerified": emailVerified,
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
