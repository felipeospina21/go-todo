package tui

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

const WIDTH = 60

type sessionState uint

const (
	defaultTime              = time.Minute
	todoView    sessionState = iota
	doneView
)

var (
	modelStyle = lipgloss.NewStyle().
			Width(WIDTH).
			Height(30).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(WIDTH).
				Height(30).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

func main() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type (
	tickMsg struct{}
	errMsg  error
)

type model struct {
	state    sessionState
	todoList list.Model
	doneList list.Model
	index    int
	err      error
}

func InitialModel() model {
	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "Bitter melon", desc: "It cools you down"},
		item{title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{title: "Eight hours of sleep", desc: "I had this once"},
		item{title: "Cats", desc: "Usually"},
	}
	done := []list.Item{
		item{title: "Plantasia, the album", desc: "My plants love it too"},
		item{title: "Pour over coffee", desc: "It takes forever to make though"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
		item{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		item{title: "Linux", desc: "Pretty much the best OS"},
		item{title: "Business school", desc: "Just kidding"},
		item{title: "Pottery", desc: "Wet clay is a great feeling"},
	}
	m := model{
		todoList: list.New(items, list.NewDefaultDelegate(), 0, 0),
		doneList: list.New(done, list.NewDefaultDelegate(), 0, 0),
		state:    todoView,
	}

	m.todoList.Title = "Todo"
	m.doneList.Title = "Done"
	// m.todoList.SetShowHelp(false)

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == todoView {
				m.state = doneView
			} else {
				m.state = todoView
			}
		}
		switch m.state {
		// update whichever model is focused
		case doneView:
			m.doneList, cmd = m.doneList.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.todoList, cmd = m.todoList.Update(msg)
			cmds = append(cmds, cmd)
		}

	case tea.WindowSizeMsg:
		h, w := docStyle.GetFrameSize()
		m.todoList.SetSize(msg.Width-w, msg.Height-h)
		m.doneList.SetSize(msg.Width-w, msg.Height-h)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var s string
	model := m.currentFocusedModel()
	if m.state == todoView {
		s += lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(fmt.Sprintf("%4s", m.todoList.View())), modelStyle.Render(m.doneList.View()))
	} else {
		s += lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(fmt.Sprintf("%4s", m.todoList.View())), focusedModelStyle.Render(m.doneList.View()))
	}
	s += helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model))
	return s
}

func (m model) currentFocusedModel() string {
	if m.state == todoView {
		return "todoList"
	}
	return "doneList"
}

// func (m *model) Next() {
// 	if m.index == len(spinners)-1 {
// 		m.index = 0
// 	} else {
// 		m.index++
// 	}
// }
