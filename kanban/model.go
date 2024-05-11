package kanban

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focused status
	loaded  bool
	lists   []list.Model
	err     error
}

func (m *Model) InitList(height int, width int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "get up", description: "get up now"},
		Task{status: todo, title: "freshen up", description: "daily routine"},
		Task{status: todo, title: "get dressed", description: "wear clothes"},
		Task{status: todo, title: "go to work", description: "get in the car and drive"},
	})

	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: inProgress, title: "cook", description: "cook food"},
		Task{status: inProgress, title: "have lunch", description: "have food"},
	})

	m.lists[complete].Title = "Complete"
	m.lists[complete].SetItems([]list.Item{
		Task{status: complete, title: "attend dsm", description: "attend standup meeting"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.InitList(msg.Height, msg.Width)
			m.loaded = true
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(lipgloss.Left,
			m.lists[todo].View(),
			m.lists[inProgress].View(),
			m.lists[complete].View(),
		)
	} else {
		return "Loading, pls wait..."
	}
}

func StartKanban() {
	m := Model{}
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Println("sth went wrong: ", err)
	}
}
