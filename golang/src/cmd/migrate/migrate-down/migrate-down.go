package migratedown

import (
	"github.com/spf13/cobra"
	"github.com/uvulpos/go-svelte/src/migrator"
)

var MigrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "🛟  Migrate your database back to a more stable version",
	Long:  `🛟  Migrate your database back to a more stable version`,
	Run: func(cmd *cobra.Command, args []string) {
		migrator.NewMigrator().MigrateDown()
	},
}

func init() {}
