package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	textarea "github.com/jbarrieault/bubble-textarea"
)

type model struct {
	width    int
	height   int
	textarea textarea.Model
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "This textarea will grow dynamically!\nTry typing multiple lines or long lines that wrap..."
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.SetPromptFunc(2, func(lineIdx int) string {
		if lineIdx == 0 {
			return "> "
		}
		return "  "
	})

	// Enable dynamic height with a maximum height
	ta.SetHeight(1)
	ta.MaxHeight = 5
	ta.SetDynamicHeight(true)

	ta.MaxWidth = 0

	ta.FocusedStyle.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Background(lipgloss.NoColor{}).
		Align(lipgloss.Center)

	ta.BlurredStyle.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Background(lipgloss.NoColor{}).
		Align(lipgloss.Center)

	return model{
		textarea: ta,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.textarea.SetWidth(min(30, m.width))
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m model) View() string {
	layout := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center)

	textArea := m.textarea.View()

	debugInfo := lipgloss.NewStyle().
		Width(m.width).
		Foreground(lipgloss.Color("240")).
		Align(lipgloss.Center).
		Render(fmt.Sprintf(
			"Visual Height: %d | Max Height: %d | "+
				"Hard Line Count: %d | Value Length: %d | m.width: %d\n",
			m.textarea.Height(),
			m.textarea.MaxHeight,
			m.textarea.LineCount(),
			len(m.textarea.Value()),
			m.width,
		))

	footer := lipgloss.NewStyle().Width(m.width).Align(lipgloss.Center).Render(textArea)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		debugInfo,
		footer,
	)

	return layout.Render(content)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
