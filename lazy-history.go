package main

import (
	"bufio"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	lib "github.com/sacenox/lazy-history/lib"
)

func main() {

	if lib.IsDebug {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			lib.Fatalf("Failed to log to file: %v", err)
		}
		defer f.Close()
	}

	query := os.Args[1] // args without program name

	lib.Debugf("Query: %s", query)

	history := []string{} // one cmd per entry

	historyFilePaths := []string{
		filepath.Join(os.Getenv("HOME"), ".history"),
		filepath.Join(os.Getenv("HOME"), ".zsh_history"),
		filepath.Join(os.Getenv("HOME"), ".bash_history"),
	}

	for _, historyFilePath := range historyFilePaths {
		if _, err := os.Stat(historyFilePath); os.IsNotExist(err) {
			lib.Debugf("History file %s does not exist", historyFilePath)
			continue
		}

		historyFile, err := os.Open(historyFilePath)
		if err != nil {
			lib.Fatalf("Failed to open history file: %v", err)
		}

		scanner := bufio.NewScanner(historyFile)
		// Read each line from the history file
		for scanner.Scan() {
			line := scanner.Text()
			history = append(history, line)
		}

		lib.Debugf("Read %d lines from history file %s", len(history), historyFilePath)

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			lib.Debugf("Error reading history file %s: %v", historyFilePath, err)
		}

		historyFile.Close()
	}

	results := lib.Search(history, query)
	p := tea.NewProgram(lib.NewHistory(results), tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		lib.Fatalf("Failed to run program: %v", err)
	}
}
