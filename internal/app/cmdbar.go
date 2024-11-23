package app

import tea "github.com/charmbracelet/bubbletea"

type cmdbar struct{}

func (b cmdbar) Init() tea.Cmd {
	return nil
}

func (b cmdbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}

func (b cmdbar) View() string {
	return "hello, cmdbar"
}
