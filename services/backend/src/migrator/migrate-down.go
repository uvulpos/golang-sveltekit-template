package migrator

import (
	"fmt"

	"github.com/uvulpos/go-svelte/src/helper/errors"
)

func (migrator Migrator) MigrateDown() *errors.Error {
	migration, migrationErr := migrator.setupMigration()
	if migrationErr != nil {
		return errors.NewInternalServerErrorApp(fmt.Sprintf("could not prepare migration setup: %s", migrationErr.Error()))
	}

	migrateErr := migration.Down()
	if migrateErr != nil && migrateErr.Error() != "no change" {
		return errors.NewInternalServerErrorApp(fmt.Sprintf("could not downgrade database: %s", migrateErr.Error()))
	}

	return nil
}
