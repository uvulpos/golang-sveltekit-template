package migrator

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
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

func (migrator Migrator) setupMigration() (*migrate.Migrate, error) {
	source, err := iofs.New(migrationDir, "migration-files")
	if err != nil {
		return nil, err
	}

	db, err := postgres.WithInstance(migrator.db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithInstance("iofs", source, "postgres", db)
	if err != nil {
		return nil, err
	}

	return m, nil
}
