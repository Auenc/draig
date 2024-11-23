package app

import (
	"github.com/auenc/draig/internal/components"
	"github.com/auenc/draig/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var topbarstyle = lipgloss.NewStyle().Background(lipgloss.Color("#00FF00"))

type topbar struct {
	heightRatio int
	methodBtn   components.Button
	test1       components.Button
	test2       components.Button
}

func (t topbar) Init() tea.Cmd {
	return nil
}

func (t topbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		topbarstyle = topbarstyle.Height(utils.ProportionalDimension(t.heightRatio, msg.Height))
		break
	}
	return t, nil
}

func (t topbar) View() string {
	//	return lipgloss.JoinHorizontal(lipgloss.Top, t.methodBtn.View(), t.test1.View(), t.test2.View())
	return topbarstyle.Render("hello, topbar")
}

func newTopbar(heightRatio int) topbar {
	return topbar{
		heightRatio: heightRatio,
		methodBtn:   components.Button{Text: "GET"},
		test1:       components.Button{Text: "test 1"},
		test2:       components.Button{Text: "test 2"},
	}
}
