package storage

import (
	"database/sql"
	"encoding/json"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

type Permission string

type UserPermissions struct {
	Permissions []Permission `db:"permissions"`
}

func (a *UserStore) GetUserPermissionsByID(tx *sqlx.Tx, userID string) ([]string, customerrors.ErrorInterface) {
	sqlquery := "SELECT get_user_permissions($1) as permissions;"

	var permissionsJSONB []byte
	var err error

	if tx != nil {
		err = tx.QueryRow(sqlquery, userID).Scan(&permissionsJSONB)
	} else {
		err = a.db.QueryRow(sqlquery, userID).Scan(&permissionsJSONB)
	}

	if err != nil {

		data := customerrors.SqlData{
			"user_id": userID,
		}

		switch err {

		case sql.ErrNoRows:
			return []string{}, customerrors.NewDatabaseNotFoundError(err, userID, sqlquery, data)

		default:
			return []string{}, customerrors.NewDatabaseError(err, userID, "Cannot get user permissions", sqlquery, data)
		}
	}

	var permissions []string
	if err := json.Unmarshal(permissionsJSONB, &permissions); err != nil {
		data := customerrors.SqlData{
			"user_id": userID,
		}
		return []string{}, customerrors.NewDatabaseError(err, userID, "Cannot unmarshal user permissions from jsonb database", sqlquery, data)
	}

	return permissions, nil
}
