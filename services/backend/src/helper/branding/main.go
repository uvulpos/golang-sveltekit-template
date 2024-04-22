package branding

import (
	"fmt"

	"github.com/uvulpos/go-svelte/src/helper/branding/models"
	"github.com/uvulpos/go-svelte/src/helper/config"
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
		Host:         config.GetWebserver().Host,
		Port:         fmt.Sprint(config.GetWebserver().Port),
		ShowFrontend: config.GetWebserver().ShowFrontend,
		ShowSwagger:  config.GetWebserver().ShowSwagger,
	}, defaultConfConfiguraton)

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println(startupPresetTable)
	fmt.Println("")
}
