/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/branding"
	webApp "github.com/uvulpos/golang-sveltekit-template/src/web-app"
)

var RunAppCmd = &cobra.Command{
	Use:   "run",
	Short: "starts the webserver to provide the application",
	Long:  `starts the webserver to provide the application`,

	Run: func(cmd *cobra.Command, args []string) {
		branding.PrintBrandingWithConfig()

		webApp.NewApp().RunApp()
	},
}
