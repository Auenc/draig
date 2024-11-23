package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	PaddingLeft(2).
	PaddingRight(2)

type Button struct {
	Text string
}

func (b Button) Init() tea.Cmd {
	return nil
}

func (b Button) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}

func (b Button) View() string {
	return style.Render(b.Text)
}
