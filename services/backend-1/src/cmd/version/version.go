/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: unknown")
	},
}
