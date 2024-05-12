package kanban

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor = 4

type Model struct {
	focused  status
	loaded   bool
	quitting bool
	lists    []list.Model
	err      error
}

func (m *Model) InitList(height int, width int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
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
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			if m.focused == todo {
				m.focused = complete
			} else {
				m.focused--
			}
		case "right", "l":
			if m.focused == complete {
				m.focused = todo
			} else {
				m.focused++
			}
		}
	case tea.WindowSizeMsg:
		if !m.loaded {
			inactiveColumnStyle.Width(msg.Width / divisor)
			focussedColumnStyle.Width(msg.Width / divisor)
			inactiveColumnStyle.Height(msg.Height - divisor)
			focussedColumnStyle.Height(msg.Height - divisor)
			m.InitList(msg.Height, msg.Width)
			m.loaded = true
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		todoView := m.lists[todo].View()
		inProgressView := m.lists[inProgress].View()
		completeView := m.lists[complete].View()

		switch m.focused {
		case complete:
			return lipgloss.JoinHorizontal(lipgloss.Left,
				inactiveColumnStyle.Render(todoView),
				inactiveColumnStyle.Render(inProgressView),
				focussedColumnStyle.Render(completeView),
			)
		case inProgress:
			return lipgloss.JoinHorizontal(lipgloss.Left,
				inactiveColumnStyle.Render(todoView),
				focussedColumnStyle.Render(inProgressView),
				inactiveColumnStyle.Render(completeView),
			)
		default:
			return lipgloss.JoinHorizontal(lipgloss.Left,
				focussedColumnStyle.Render(todoView),
				inactiveColumnStyle.Render(inProgressView),
				inactiveColumnStyle.Render(completeView),
			)
		}
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
