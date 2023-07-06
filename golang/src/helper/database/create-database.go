package database

import (
	"database/sql"
	"fmt"
)

func CreateDatabase() *sql.DB {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		"postgres",
		"postgres",
		"postgres:5432",
		"postgres",
		"disable",
	)
	db, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		panic(dbErr)
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
