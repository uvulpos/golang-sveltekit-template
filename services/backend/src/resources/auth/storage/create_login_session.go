package storage

import (
	"database/sql"
	"fmt"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func (a *AuthStore) StartLoginSession(tx *sqlx.Tx, loggedinUser string, useragentHash, ipaddr *string) (string, customerrors.ErrorInterface) {
	sqlquery := "INSERT INTO user_sessions (user_id, useragent_hash, created_ip_addr, last_jwt_refresh_ip_addr) VALUES ($1, $2, $3, $4) RETURNING id"

	data := customerrors.SqlData{
		"user_id":                  loggedinUser,
		"useragent_hash":           useragentHash,
		"created_ip_addr":          ipaddr,
		"last_jwt_refresh_ip_addr": ipaddr,
	}

	fmt.Println(sqlquery)
	fmt.Println(data)

	var sessionID string
	err := tx.Get(&sessionID, sqlquery, loggedinUser, useragentHash, ipaddr, ipaddr)
	if err != nil {

		data := customerrors.SqlData{
			"user_id":                  loggedinUser,
			"useragent_hash":           useragentHash,
			"created_ip_addr":          ipaddr,
			"last_jwt_refresh_ip_addr": ipaddr,
		}

		switch err {

		case sql.ErrNoRows:
			return "", customerrors.NewDatabaseNotFoundError(err, "", sqlquery, data)

		default:
			return "", customerrors.NewDatabaseError(err, "", "Cannot create User Session", sqlquery, data)
		}
	}

	return sessionID, nil
}
