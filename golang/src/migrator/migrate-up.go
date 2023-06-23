package migrator

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func (migrator Migrator) Migrate() {
	source, err := iofs.New(migrationDir, "migration-files")
	if err != nil {
		panic(err)
	}

	db, err := postgres.WithInstance(migrator.db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database driver: %v", err)
	}

	m, err := migrate.NewWithInstance("iofs", source, "postgres", db)
	if err != nil {
		panic(err)
	}

	// err = m.Down()
	err = m.Up()
	if err != nil {
		panic(err)
	}

	fmt.Println("Do Migration stuff")
}
