package database

import (
	"fmt"

	"github.com/go-sqlx/sqlx"
	"github.com/uvulpos/go-svelte/src/helper/config"
)

func CreateDatabase(configuration *config.Configuration) Sql {
	dbHost := configuration.DB.Host
	dbPort := configuration.DB.Port
	dbUsername := configuration.DB.Username
	dbPassword := configuration.DB.Password
	dbDatabase := configuration.DB.Database

	addr := fmt.Sprintf("%s:%d", dbHost, dbPort)

	var sslMode string = "disable"
	if configuration.DB.SslMode {
		sslMode = "require"
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		dbUsername,
		dbPassword,
		addr,
		dbDatabase,
		sslMode,
	)

	fmt.Println("DB CONN: ", connStr)
	db, dbErr := sqlx.Connect("postgres", connStr)
	// db, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		panic(dbErr)
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	var sqlWrapper Sql = Sql{
		DB: db,
	}

	return sqlWrapper
}
