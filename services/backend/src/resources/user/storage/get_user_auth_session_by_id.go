package storage

import (
	"database/sql"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"

	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
	storageModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/storage/models"
)

func (s *UserStore) GetUserAuthSessionByID(tx *sqlx.Tx, sessionID string) (*serviceModel.SessionModel, customerrors.ErrorInterface) {
	sqlquery := "SELECT * FROM user_sessions WHERE id=$1 LIMIT 1;"
	var session storageModel.SessionModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, sessionID).StructScan(&session)
	} else {
		err = s.db.QueryRowx(sqlquery, sessionID).StructScan(&session)
	}

	if err != nil {
		data := customerrors.SqlData{
			"sessionID": sessionID,
		}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, "", "Cannot get session", sqlquery, data)
		}
	}

	return session.ToServiceModel(), nil
}
