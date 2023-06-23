/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	webApp "github.com/uvulpos/go-svelte/src/web-app"
)

var (
	showNoFrontend bool
	webserverPort  int
)

var RunAppCmd = &cobra.Command{
	Use:   "run",
	Short: "starts the webserver to provide the application",
	Long:  `starts the webserver to provide the application`,
	Run: func(cmd *cobra.Command, args []string) {
		webApp.NewApp().RunApp(
			!showNoFrontend,
			webserverPort,
		)
	},
}

func init() {
	RunAppCmd.Flags().BoolVar(&showNoFrontend, "no-frontend", false, "disable the frontend and deliver just an api")
	RunAppCmd.Flags().IntVarP(&webserverPort, "port", "p", 3000, "define a port, where the application should listen to")
}
