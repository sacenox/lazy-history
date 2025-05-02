package lib

import (
	"bytes"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type History struct {
	Entries  []string
	Selected int
	Output   *bytes.Buffer
}

type ExecuteEntryMsg struct {
	Error error
}

func NewHistory(entries []string) *History {
	return &History{
		Entries:  entries,
		Selected: 0,
		Output:   &bytes.Buffer{},
	}
}

func (h *History) Init() tea.Cmd {
	return nil
}

func (h *History) ExecuteEntry(entry string) tea.Cmd {
	Debugf("executing: %s", entry)
	parts := strings.Split(entry, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = h.Output
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		Debugf("error: %v", err)
		return ExecuteEntryMsg{Error: err}
	})
}

func (h *History) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case ExecuteEntryMsg:
		Debugf("ExecuteEntryMsg: %v", msg)
		if msg.Error != nil {
			Debugf("error: %v", msg.Error)
			return h, nil
		}
		return h, tea.Quit

	// Is it a key press?
	case tea.KeyMsg:
		Debugf("key: %s", msg.String())

		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return h, tea.Quit
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if h.Selected > 0 {
				h.Selected--
			}
		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if h.Selected < len(h.Entries)-1 {
				h.Selected++
			}
		case "enter":
			// return the selected entry, and run the command
			Debugf("executing: %s", h.Entries[h.Selected])
			return h, h.ExecuteEntry(h.Entries[h.Selected])
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return h, nil
}

func (h *History) View() string {
	// render the history
	output := h.Output.String()
	if output != "" {
		return output
	}

	history := ""

	// Print all entries with the selected one highlighted
	for i, entry := range h.Entries {
		if i == h.Selected {
			history += "> " + entry + "\n"
		} else {
			history += "  " + entry + "\n"
		}
	}

	return history
}
