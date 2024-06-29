package database

import (
	"fmt"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/configuration"
)

func CreateDatabase() (*Sql, error) {
	connStr := GetSqlConnectionString()

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

func GetSqlConnectionString() string {

	var databaseConfig = configuration.Configuration.Database

	var sslMode string = "disable"
	if databaseConfig.SSL {
		sslMode = "require"
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Addr,
		databaseConfig.Port,
		databaseConfig.Database,
		sslMode,
	)

	// humanString := fmt.Sprintf(
	// 	"%s:%d",
	// 	databaseConfig.Addr,
	// 	databaseConfig.Port,
	// )

	return connString
}
