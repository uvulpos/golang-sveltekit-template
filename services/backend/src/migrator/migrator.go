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

//go:embed migration-files-app/*.sql
var migrationAppDir embed.FS

// //go:embed migration-files-test/*.sql
// var migrationTestDir embed.FS

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

func (m *Migrator) MigrateUp() error {
	return m.migrateDatabase("app")
}

func (m *Migrator) migrateDatabase(migrationPrefix string) error {
	driver, err := postgres.WithInstance(m.db, new(postgres.Config))
	if err != nil {
		return err
	}

	var migrationDirectory embed.FS = migrationAppDir

	sourceInstance, err := httpfs.New(http.FS(migrationDirectory), fmt.Sprintf("migration-files-%s", migrationPrefix))
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
