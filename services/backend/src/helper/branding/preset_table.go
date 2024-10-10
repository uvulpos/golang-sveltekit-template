package branding

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/branding/models"
)

func getAsciiConfigurationTable(configuration models.ConfigurationTable, presetStle lipgloss.Style) string {
	configurationRows := [][]string{
		{"URL:", configuration.DisplayHost},
		{"", ""},
		{"Host (Container) :", configuration.Host},
		{"Port (Container) :", configuration.Port},
		{"", ""},
		{"ShowFrontend:", fmt.Sprintf("%t", configuration.ShowFrontend)},
		{"ShowSwagger:", fmt.Sprintf("%t", configuration.ShowSwagger)},
	}

	informationTable := table.New().
		Border(lipgloss.NormalBorder()).
		Headers("Key", "Value").
		Width(90).
		Rows(configurationRows...)

	styleInformation := presetStle.Align(lipgloss.Left).
		PaddingLeft(3).
		PaddingRight(3).
		Underline(false).
		Bold(false)

	return styleInformation.Render(informationTable.Render())
}
