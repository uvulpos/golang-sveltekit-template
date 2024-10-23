package storage

import (
	"database/sql"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
	serviceModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/service/models"
	storageModel "github.com/uvulpos/golang-sveltekit-template/src/resources/user/storage/models"
)

func (s *UserStore) GetUserByUsername(tx *sqlx.Tx, username string) (*serviceModel.UserModel, customerrors.ErrorInterface) {

	sqlquery := "SELECT * FROM users WHERE username=$1 LIMIT 1;"
	var user storageModel.UserModel
	var err error

	if tx != nil {
		err = tx.QueryRowx(sqlquery, username).StructScan(&user)
	} else {
		err = s.db.QueryRowx(sqlquery, username).StructScan(&user)
	}

	if err != nil {
		data := customerrors.SqlData{
			"username": username,
		}

		switch err {

		case sql.ErrNoRows:
			return nil, customerrors.NewDatabaseNotFoundError(err, username, sqlquery, data)

		default:
			return nil, customerrors.NewDatabaseError(err, username, "Cannot get user", sqlquery, data)
		}
	}

	return user.ToServiceModel(), nil
}
