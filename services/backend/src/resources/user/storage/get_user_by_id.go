package storage

import (
	"database/sql"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
	storageModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/storage/models"
)

func (s *UserStore) GetUserByID(tx *sqlx.Tx, userID string) (*serviceModel.UserModel, customerrors.ErrorInterface) {

	sqlquery := "SELECT * FROM users WHERE id=$1 LIMIT 1;"
	var user storageModel.UserModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, userID).StructScan(&user)
	} else {
		err = s.db.QueryRowx(sqlquery, userID).StructScan(&user)
	}

	if err != nil {
		data := customerrors.SqlData{
			"user_id": userID,
		}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, userID, sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, userID, "Cannot get user", sqlquery, data)
		}
	}

	return user.ToServiceModel(), nil
}
