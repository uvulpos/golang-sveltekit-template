package migrator

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/lib/pq"
	"github.com/uvulpos/go-svelte/src/helper/config"
	dbHelper "github.com/uvulpos/go-svelte/src/helper/database"
)

//go:embed migration-files/*.sql
var migrationDir embed.FS

type Migrator struct {
	db *sql.DB
}

func NewMigrator(configuration *config.Configuration) *Migrator {
	return &Migrator{
		db: dbHelper.CreateDatabase(configuration).DB.DB,
	}
}

func (m Migrator) MigrateUp() error {
	driver, err := postgres.WithInstance(m.db, new(postgres.Config))
	if err != nil {
		return err
	}
	sourceInstance, err := httpfs.New(http.FS(migrationDir), "migration-files")
	if err != nil {
		return err
	}
	migrator, migratorErr := migrate.NewWithInstance("httpfs", sourceInstance, "postgres", driver)
	if migratorErr != nil {
		return migratorErr
	}
	err = migrator.Up()
	return err
}
