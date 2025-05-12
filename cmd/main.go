package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sacenox/lazy-history/internal/debug"
	internal_list "github.com/sacenox/lazy-history/internal/list"
	"github.com/sacenox/lazy-history/internal/search"
)

func main() {
	if debug.IsDebug {
		// Clear the debug.log file
		os.Truncate("debug.log", 0)

		// Log to the debug.log file
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatalf("Failed to log to file: %v", err)
		}
		defer f.Close()
	}

	// Get the output of `history`
	history := []string{} // one cmd per entry

	historyFilePaths := []string{
		filepath.Join(os.Getenv("HOME"), ".history"),
		filepath.Join(os.Getenv("HOME"), ".zsh_history"),
		filepath.Join(os.Getenv("HOME"), ".bash_history"),
	}

	// TODO: This is not working, we need to do it in an interactive shell owner by the user.
	// Run `history -a` to append the history to the files
	exec.Command("history", "-a").Run()

	// Read the history files
	for _, historyFilePath := range historyFilePaths {
		if _, err := os.Stat(historyFilePath); os.IsNotExist(err) {
			continue
		}

		historyFile, err := os.Open(historyFilePath)
		if err != nil {
			log.Fatalf("Failed to open history file: %v", err)
		}

		scanner := bufio.NewScanner(historyFile)
		// Read each line from the history file
		for scanner.Scan() {
			line := scanner.Text()
			history = append(history, line)
		}

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading history file %s: %v", historyFilePath, err)
		}

		historyFile.Close()
	}

	debug.Debugf("History length: %d", len(history))

	// Join all args into a single string
	query := strings.Join(os.Args[1:], " ")

	debug.Debugf("Searching for %s", query)

	// Search the history output for the query
	results := search.Search(history, query)

	debug.Debugf("Found %d results", len(results))
	debug.Debugf("Results: %v", results)

	ui := internal_list.New(results)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		log.Fatal("Error running program: ", err)
	}
}
