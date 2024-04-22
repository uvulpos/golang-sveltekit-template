/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uvulpos/go-svelte/src/helper/branding"
	"github.com/uvulpos/go-svelte/src/helper/config"
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
		config.LoadData()
		branding.PrintBrandingWithConfig()

		webApp.NewApp().RunApp()
	},
}
