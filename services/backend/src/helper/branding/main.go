package branding

import (
	"fmt"

	"github.com/uvulpos/go-svelte/src/helper/branding/models"
)

func PrintBranding() {
	logo, logoErr := getAsciiLogo()
	if logoErr != nil {
		panic(logoErr)
	}

	startupInformation, _ := getAsciiStartupInformation()

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println("")
}

func PrintBrandingWithConfig() {
	logo, logoErr := getAsciiLogo()
	if logoErr != nil {
		panic(logoErr)
	}

	startupInformation, defaultConfConfiguraton := getAsciiStartupInformation()

	startupPresetTable := getAsciiConfigurationTable(models.ConfigurationTable{
		Host: "127.0.0.1",
		Port: "8080",
	}, defaultConfConfiguraton)

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println(startupPresetTable)
	fmt.Println("")
}
