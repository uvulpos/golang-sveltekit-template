package migrator

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
)

//go:embed migration-files/*.sql
var migrationDir embed.FS

type Migrator struct {
	db *sql.DB
}

func NewMigrator() *Migrator {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		"postgres",
		"mysecretpassword",
		"127.0.0.1:5432",
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

	return &Migrator{
		db,
	}
}
