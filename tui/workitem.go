package tui

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"log"
)

func RunWorkItem() {
	m := workItemModel{}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Printf("[RunWorkItem] error: %v", err)
	}
}

type workItemModel struct {}

func (m workItemModel) Init() tea.Cmd {
	return nil
}

func (m workItemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m workItemModel) View() string {
	return "Work Item TUI (CRUD operations coming soon)\nPress Ctrl+C to exit."
}
