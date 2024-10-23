package storage

import (
	"encoding/json"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
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

		default:
			return []string{}, customerrors.NewDatabaseError(err, userID, "Cannot get user permissions", sqlquery, data)
		}
	}

	// if result is nil, user has no role or permissions yet
	if permissionsJSONB == nil {
		return []string{}, nil
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
