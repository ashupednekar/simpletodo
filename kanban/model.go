package kanban

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list list.Model
	err  error
}

func (m *Model) InitList(height int, width int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "TODO"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "get up", description: "get up now"},
		Task{status: todo, title: "freshen up", description: "daily routine"},
		Task{status: todo, title: "get dressed", description: "wear clothes"},
		Task{status: todo, title: "go to work", description: "get in the car and drive"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.InitList(msg.Height, msg.Width)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}

func StartKanban() {
	m := Model{}
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Println("sth went wrong: ", err)
	}
}
