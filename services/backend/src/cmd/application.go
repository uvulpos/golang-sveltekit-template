package cmd

import (
	"log"

	"github.com/spf13/cobra"

	migrate "github.com/uvulpos/go-svelte/src/cmd/migrate"
	runApp "github.com/uvulpos/go-svelte/src/cmd/run"
	version "github.com/uvulpos/go-svelte/src/cmd/version"
	"github.com/uvulpos/go-svelte/src/helper/config"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "application",
	Short: "awesome application, build with golang + svelte",
	Long:  `awesome application, build with golang + svelte`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	runCmd.AddCommand(runApp.RunAppCmd)
	runCmd.AddCommand(migrate.MigrateCmd)
	runCmd.AddCommand(version.VersionCmd)
}

func Execute() {
	config.LoadData()
	err := runCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
