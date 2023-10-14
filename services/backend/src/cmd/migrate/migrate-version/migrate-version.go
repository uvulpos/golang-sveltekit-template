package migrateversion

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uvulpos/go-svelte/src/migrator"

	"github.com/fatih/color"
)

var (
	schemaMigrateVersion uint
	forceMigration       bool
)

var MigrateUpCmd = &cobra.Command{
	Use:   "version",
	Short: "🚀 Migrate your database to a specific version",
	Long:  `🚀 Migrate your database to a specific version`,
	Run: func(cmd *cobra.Command, args []string) {
		if schemaMigrateVersion != 999999999 {
			err := migrator.NewMigrator().MigrateVersion(schemaMigrateVersion, forceMigration)
			if err != nil {
				panic(err)
			}
		} else {
			c := color.New(color.FgRed).Add(color.Bold)
			c.Println("🚨 PLEASE SPECIFY A VERSION!")
			fmt.Println()
			cmd.Help()
		}
	},
}

func init() {
	MigrateUpCmd.Flags().UintVarP(&schemaMigrateVersion, "schema-version", "v", 999999999, "Define a schema version you want to migrate to")
	MigrateUpCmd.Flags().BoolVarP(&forceMigration, "force", "f", false, "Force migration if database is corrupt")
}
