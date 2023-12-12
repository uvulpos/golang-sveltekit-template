package migrateup

import (
	"github.com/spf13/cobra"
	"github.com/uvulpos/go-svelte/src/migrator"
)

var MigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "🚀 Migrate your database to a newer version",
	Long:  `🚀 Migrate your database to a newer version`,
	Run: func(cmd *cobra.Command, args []string) {
		err := migrator.NewMigrator().MigrateUp()
		if err != nil {
			panic(err)
		}
	},
}

func init() {}
