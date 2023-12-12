package database

import (
	"fmt"

	"github.com/go-sqlx/sqlx"
)

func CreateDatabase() Sql {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		"postgres",
		"postgres",
		"postgres:5432",
		"postgres",
		"disable",
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
