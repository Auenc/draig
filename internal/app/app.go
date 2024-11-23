package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focus int

const (
	topBarIndex       = 0
	method      focus = iota
	url
	request
	response
)

var style = lipgloss.NewStyle().PaddingLeft(2).PaddingRight(2).PaddingTop(1)

type model struct {
	topbar      topbar
	requestbar  requestbar
	responsebar responsebar
	cmdbar      cmdbar
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Draig")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.topbar.Update(msg)
		break
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	views := []string{
		m.topbar.View(),
		m.requestbar.View(),
		m.responsebar.View(),
		m.cmdbar.View(),
	}
	joinedView := lipgloss.JoinVertical(lipgloss.Left, views...)

	return style.Render(joinedView)
}

func initialModel() model {
	m := model{topbar: newTopbar(2)}

	return m
}

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong %v", err)
		os.Exit(1)
	}
}
