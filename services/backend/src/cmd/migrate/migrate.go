/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	gomigrator "github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/branding"
	"github.com/uvulpos/golang-sveltekit-template/src/migrator"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate-db",
	Short: "🚀 Migrate your database to a newer version",
	Long:  `🚀 Migrate your database to a newer version`,
	Run: func(cmd *cobra.Command, args []string) {

		branding.PrintBranding()
		err := migrator.NewMigrator().MigrateUp()

		messageStyle := lipgloss.NewStyle().Bold(true)
		successMessageStyle := messageStyle.Foreground(lipgloss.Color("#1eb523"))
		errorMessageStyle := messageStyle.Foreground(lipgloss.Color("#c42f2d"))

		if err != nil {

			switch err {
			case gomigrator.ErrNoChange:
				fmt.Println(successMessageStyle.Render("Database is on the newest version"))
			case gomigrator.ErrNilVersion:
				_ = fmt.Errorf("%s", errorMessageStyle.Render("No Migration found."))
			case gomigrator.ErrInvalidVersion:
				_ = fmt.Errorf("%s", errorMessageStyle.Render("Database is on a newer version than your software. Please use the newer version"))
			case gomigrator.ErrLocked:
				_ = fmt.Errorf("%s", errorMessageStyle.Render("Database is currently locked by another application"))
			case gomigrator.ErrLockTimeout:
				_ = fmt.Errorf("%s", errorMessageStyle.Render("Database Migration timed out"))
			default:
				panic(err)
			}

			os.Exit(0)
		}

		fmt.Println(successMessageStyle.Render("Database Migration was successful"))
		os.Exit(0)
	},
}

func init() {}
