package cmd

import (
	"log"
	// "time"

	"github.com/spf13/cobra"

	// "github.com/spf13/cobra/doc"
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
	// header := &doc.GenManHeader{
	// 	Title:   "MINE",
	// 	Section: "3",
	// }
	// err := doc.GenManTree(runCmd, header, "/tmp")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// time.Sleep(time.Minute * 10)

	config.LoadData()
	err := runCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
