package kanban

import "github.com/charmbracelet/lipgloss"

var (
	inactiveColumnStyle = lipgloss.NewStyle().
				Padding(1, 2)
	focussedColumnStyle = lipgloss.NewStyle().
				Padding(1, 2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
