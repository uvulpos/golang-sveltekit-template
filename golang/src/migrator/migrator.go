package migrator

import (
	"database/sql"
	"embed"

	_ "github.com/lib/pq"
	dbHelper "github.com/uvulpos/go-svelte/src/helper/database"
)

//go:embed migration-files/*.sql
var migrationDir embed.FS

type Migrator struct {
	db *sql.DB
}

func NewMigrator() *Migrator {
	return &Migrator{
		db: dbHelper.CreateDatabase(),
	}
}
