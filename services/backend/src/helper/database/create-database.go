package database

import (
	"fmt"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

func CreateDatabase() (*sqlx.DB, error) {
	connStr := GetSqlConnectionString()

	db, dbErr := sqlx.Connect("postgres", connStr)
	if dbErr != nil {
		return nil, dbErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}

func GetSqlConnectionString() string {

	var sslMode string = "disable"
	if configuration.DATABASE_SSL {
		sslMode = "require"
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		configuration.DATABASE_USERNAME,
		configuration.DATABASE_PASSWORD,
		configuration.DATABASE_ADDR,
		configuration.DATABASE_PORT,
		configuration.DATABASE_DATABASE,
		sslMode,
	)

	return connString
}
