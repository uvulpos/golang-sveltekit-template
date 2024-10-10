package setup

import (
	"github.com/uvulpos/golang-sveltekit-template/src/migrator"
)

func MigrateDatabaseWithTestData() error {
	migratorInst := migrator.NewMigrator()
	return migratorInst.MigrateTestData()
}
