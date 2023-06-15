package migrator

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (migrator Migrator) Migrate() {

	// create postgres connection via env variables or flags
	// create backup first, before start migrating

	// source, err := iofs.New(migrationDir, "migration-files")
	// if err != nil {
	// 	panic(err)
	// }
	// m, err := migrate.NewWithSourceAndDatabaseInstance("iofs", source, migrator.db)
	// if err != nil {
	// 	panic(err)
	// }

	// // err = m.Down()
	// err = m.Up()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Do Migration stuff")
}
