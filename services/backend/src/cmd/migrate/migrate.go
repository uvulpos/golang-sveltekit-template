/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	migratedown "github.com/uvulpos/go-svelte/src/cmd/migrate/migrate-down"
	migrateup "github.com/uvulpos/go-svelte/src/cmd/migrate/migrate-up"
	migrateversion "github.com/uvulpos/go-svelte/src/cmd/migrate/migrate-version"
)

// migrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate-db",
	Short: "Migrate your database to the newest/highest or a stable version",
	Long: `Migrate your database to the newest/highest or a stable version. 

🚨 This command is for experts only who know what they are doing 🚨

This makes sense, when you just updated the application via a package manager 
and you need to migrate everything to the newest version. If you notice an error 
in the application, check your version and specify the database migration ID, 
which you can find in the release notes.
`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	MigrateCmd.AddCommand(migrateup.MigrateUpCmd)
	MigrateCmd.AddCommand(migratedown.MigrateDownCmd)
	MigrateCmd.AddCommand(migrateversion.MigrateUpCmd)
}
