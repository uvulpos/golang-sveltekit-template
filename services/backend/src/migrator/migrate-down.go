package migrator

func (migrator Migrator) MigrateDown() {
	migration, migrationErr := migrator.setupMigration()
	if migrationErr != nil {
		panic(migrationErr)
	}

	migrateErr := migration.Down()
	if migrateErr != nil && migrateErr.Error() != "no change" {
		panic(migrateErr)
	}
}
