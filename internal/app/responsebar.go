package app

import tea "github.com/charmbracelet/bubbletea"

type responsebar struct{}

func (r responsebar) Init() tea.Cmd {
	return nil
}

func (r responsebar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

func (r responsebar) View() string {
	return "hello, response bar\n"
}
