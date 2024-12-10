package components

import (
	"github.com/auenc/draig/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Container struct {
	style       lipgloss.Style
	heightRatio float64
	children    []tea.Model
}

func (c Container) Init() tea.Cmd {
	return nil
}

func (c *Container) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		height := utils.ProportionalDimension(c.heightRatio, msg.Height)
		c.style = c.style.Height(height).Width(msg.Width)
		break
	}
	return c, nil
}

func (c Container) View() string {
	children := []string{}
	for _, child := range c.children {
		children = append(children, child.View())
	}
	content := lipgloss.JoinHorizontal(lipgloss.Top, children...)
	return c.style.Render(content)
}

func NewContainer(heightRatio float64, children []tea.Model) Container {
	return Container{
		style:       lipgloss.NewStyle(),
		heightRatio: heightRatio,
		children:    children,
	}
}
