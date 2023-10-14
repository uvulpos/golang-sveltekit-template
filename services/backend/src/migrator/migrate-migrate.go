package migrator

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uvulpos/go-svelte/src/helper/errors"
)

func (migrator Migrator) MigrateVersion(version uint, force bool) *errors.Error {
	migration, migrationErr := migrator.setupMigration()
	if migrationErr != nil {
		return errors.NewInternalServerErrorApp(fmt.Sprintf("could not prepare migration setup: %s", migrationErr.Error()))
	}

	if force {
		migrateErr := migration.Force(int(version))
		if migrateErr != nil && migrateErr.Error() != "no change" {
			return errors.NewInternalServerErrorApp(fmt.Sprintf("could not execute the forced migration: %s", migrateErr.Error()))
		}

	} else {
		migrateErr := migration.Migrate(version)
		if migrateErr != nil && migrateErr.Error() != "no change" {
			return errors.NewInternalServerErrorApp(fmt.Sprintf("could not execute the migration: %s", migrateErr.Error()))
		}
	}

	return nil
}
