package app

import (
	"fmt"
	"os"

	"github.com/auenc/draig/internal/components"
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

var style = lipgloss.NewStyle() //.PaddingLeft(2).PaddingRight(2).PaddingTop(1)

type model struct {
	topbar      components.Container
	requestbar  components.Container
	responsebar components.Container
	cmdbar      components.Container
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Draig")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.topbar.Update(msg)
		m.requestbar.Update(msg)
		m.responsebar.Update(msg)
		m.cmdbar.Update(msg)
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
	m := model{
		topbar: components.NewContainer(1, []tea.Model{
			components.Button{Text: "GET"},
			components.Button{Text: "test 1"},
			components.Button{Text: "test 2"},
		}),
		requestbar:  components.NewContainer(4, nil),
		responsebar: components.NewContainer(6.5, nil),
		cmdbar:      components.NewContainer(0.5, nil),
	}

	return m
}

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong %v", err)
		os.Exit(1)
	}
}
