package migrator

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/lib/pq"
	dbHelper "github.com/uvulpos/golang-sveltekit-template/src/helper/database"
)

//go:embed migration-files/*.sql
var migrationDir embed.FS

type Migrator struct {
	db *sql.DB
}

func NewMigrator() *Migrator {
	dbConn, dbConnErr := dbHelper.CreateDatabase()
	if dbConn == nil || dbConn.DB == nil || dbConnErr != nil {
		err := fmt.Errorf("could not connect to database")
		if err != nil {
			panic(err)
		}
	}
	return &Migrator{
		db: (*dbConn).DB,
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
