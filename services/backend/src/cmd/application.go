package cmd

import (
	"log"
	"os"

	// "time"

	"github.com/spf13/cobra"

	// "github.com/spf13/cobra/doc"
	migrate "github.com/uvulpos/go-svelte/src/cmd/migrate"
	runApp "github.com/uvulpos/go-svelte/src/cmd/run"
	"github.com/uvulpos/go-svelte/src/helper/branding"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "application",
	Short: "awesome application, build with golang + svelte",
	Long:  `awesome application, build with golang + svelte`,
	Run: func(cmd *cobra.Command, args []string) {
		branding.PrintBranding()
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	runCmd.AddCommand(runApp.RunAppCmd)
	runCmd.AddCommand(migrate.MigrateCmd)
	// runCmd.AddCommand(version.VersionCmd)
}

func Execute() {
	err := runCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
