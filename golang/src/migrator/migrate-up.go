package migrator

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (migrator Migrator) Migrate() {
	migration, migrationErr := migrator.setupMigration()
	if migrationErr != nil {
		panic(migrationErr)
	}

	migrateErr := migration.Up()
	if migrateErr != nil && migrateErr.Error() != "no change" {
		panic(migrateErr)
	}
}
