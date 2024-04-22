package database

import (
	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/helper/config"
)

func CreateDatabase() (*Sql, error) {
	connStr, _ := config.GetSqlConnectionString()

	db, dbErr := sqlx.Connect("postgres", connStr)
	if dbErr != nil {
		return nil, dbErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return &Sql{
		DB: db,
	}, nil
}
