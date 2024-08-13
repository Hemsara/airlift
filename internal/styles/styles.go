package styles

import "github.com/charmbracelet/lipgloss"

var ErrStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF0000")).
	Background(lipgloss.Color("#000000"))

var SuccessStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#000000")).
	PaddingLeft(4).
	PaddingRight(4).
	MarginBottom(2).
	Background(lipgloss.Color("#FFFFFF"))
