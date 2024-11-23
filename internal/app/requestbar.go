package app

import tea "github.com/charmbracelet/bubbletea"

type requestbar struct{}

func (r requestbar) Init() tea.Cmd {
	return nil
}

func (r requestbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

func (r requestbar) View() string {
	return "hello, request bar\n"
}
