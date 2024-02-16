package branding

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/uvulpos/go-svelte/src/helper/branding/models"
)

func getAsciiConfigurationTable(configuration models.ConfigurationTable, presetStle lipgloss.Style) string {
	configurationRows := [][]string{
		{"Host:", configuration.Host},
		{"Port:", configuration.Port},
	}

	informationTable := table.New().
		Border(lipgloss.NormalBorder()).
		Headers("Key", "Value").
		Width(90).
		Rows(configurationRows...)

	styleInformation := presetStle.Align(lipgloss.Left).
		PaddingLeft(5).
		PaddingRight(5).
		Underline(false).
		Bold(false)

	return styleInformation.Render(informationTable.Render())
}
