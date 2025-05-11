package internal_list

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

type item string

func (i item) FilterValue() string {
	return ""
}

type itemDelegate struct{}

func (d itemDelegate) Height() int {
	return 1
}

func (d itemDelegate) Spacing() int {
	return 0
}

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	if index == m.Index() {
		fmt.Fprint(w, "→ ")
	}

	fmt.Fprint(w, i)
}

type model struct {
	list   list.Model
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return ""
	}

	return "\n" + m.list.View()
}

func New(lines []string) model {
	items := lo.Map(lines, func(line string, _ int) list.Item {
		return item(line)
	})

	highlight := "13"
	height := 20
	width := 80

	l := list.New(items, itemDelegate{}, width, height)
	l.Title = "—  Lazy History —"
	l.Styles.Title = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(highlight))
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)

	return model{
		list: l,
	}
}
