package migrator

import (
	"fmt"
	"regexp"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uvulpos/go-svelte/src/helper/errors"
)

const doesNotExistPattern = `(?m)no migration found for version [0-9]+: read down for version [0-9]+ migration-files: file does not exist`

func (migrator Migrator) MigrateUp() *errors.Error {
	migration, migrationErr := migrator.setupMigration()
	if migrationErr != nil {
		return errors.NewInternalServerErrorApp(fmt.Sprintf("could not prepare migration setup: %s", migrationErr.Error()))
	}

	migrateErr := migration.Up()
	if migrateErr != nil {

		// no newer migration found (package says it's an error, but isn't, but should be communicated to the user)
		if match, matchErr := regexp.Match(doesNotExistPattern, []byte(migrateErr.Error())); matchErr == nil && match {
			fmt.Println("🚀 database is up2date, no migration needed")
			return nil
		}

		return errors.NewInternalServerErrorApp(fmt.Sprintf("could not execute the up migration: %s", migrateErr.Error()))
	}

	return nil
}
