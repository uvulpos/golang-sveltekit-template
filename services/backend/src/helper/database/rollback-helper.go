package database

import "github.com/go-sqlx/sqlx"

func Rollback(tx *sqlx.Tx) {
	_ = tx.Rollback()
}
