package lib

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListItem struct {
	title string
}

func (l ListItem) Title() string {
	return l.title
}

func (l ListItem) Description() string {
	return ""
}

func (l ListItem) FilterValue() string {
	return l.title
}

type History struct {
	Entries  []string
	Selected int
	Output   *bytes.Buffer
	List     list.Model
}

type ExecuteEntryMsg struct {
	Error error
}

func NewHistory(entries []string) *History {
	items := make([]list.Item, len(entries))
	for i, entry := range entries {
		items[i] = ListItem{title: entry}
	}
	return &History{
		Selected: 0,
		Output:   &bytes.Buffer{},
		List:     list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}

func (h *History) Init() tea.Cmd {
	return nil
}

func (h *History) ExecuteEntry(entry string) tea.Cmd {
	parts := strings.Split(entry, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = h.Output
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return ExecuteEntryMsg{Error: err}
	})
}

func (h *History) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case ExecuteEntryMsg:
		if msg.Error != nil {
			Debugf("error: %v", msg.Error)
			return h, nil
		}
		return h, tea.Quit

	case tea.WindowSizeMsg:
		h.List.SetSize(msg.Width, msg.Height)

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return h, tea.Quit
		case "enter":
			// return the selected entry, and run the command
			selected := h.List.SelectedItem().(ListItem)
			Debugf("executing: %s", selected.title)
			return h, h.ExecuteEntry(selected.title)
		}
	}

	var cmd tea.Cmd
	h.List, cmd = h.List.Update(msg)
	return h, cmd
}

func (h *History) View() string {
	// render the history
	output := h.Output.String()
	if output != "" {
		return output
	}

	return h.List.View()
}
