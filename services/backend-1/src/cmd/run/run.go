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
	showSwaggger   bool
	webserverPort  int
)

var RunAppCmd = &cobra.Command{
	Use:   "run",
	Short: "starts the webserver to provide the application",
	Long:  `starts the webserver to provide the application`,

	Run: func(cmd *cobra.Command, args []string) {
		webApp.NewApp().RunApp(
			!showNoFrontend,
			showSwaggger,
			webserverPort,
		)
	},
}

func init() {
	RunAppCmd.Flags().BoolVar(&showNoFrontend, "no-frontend", false, "disable the frontend and deliver just an api")
	RunAppCmd.Flags().BoolVar(&showSwaggger, "show-swagger", false, "enable the swagger api for developing")
	RunAppCmd.Flags().IntVarP(&webserverPort, "port", "p", 3000, "define a port, where the application should listen to")
}