package cli

import "github.com/charmbracelet/lipgloss"

var Help_Style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#00b7ff"))

var Red_Style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#812e23"))

var Info_Style = lipgloss.NewStyle().
	SetString("INFO").
	Bold(true).
	Foreground(lipgloss.Color("#00ff37")).
	Background(lipgloss.Color("#008643"))

var Error_Style = lipgloss.NewStyle().
	SetString("ERROR").
	Bold(true).
	Foreground(lipgloss.Color("#ff0055")).
	Background(lipgloss.Color("#3d0014"))
