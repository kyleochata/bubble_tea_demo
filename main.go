package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model_cmd struct {
	status int
	err    error
}
type statusMsg int
type errMsg struct{ err error }

func checkServer() tea.Msg {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		return errMsg{err}
	}
	return statusMsg(response.StatusCode)
}
func (e errMsg) Error() string { return e.err.Error() }
func (m model_cmd) Init() tea.Cmd {
	return checkServer
}
func (m model_cmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m model_cmd) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWE had some trouble: %v\n\n", m.err)
	}
	s := fmt.Sprintf("Checking %s...", url)
	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}
	return "\n" + s + "\n\n"
}

func main() {
	if _, err := tea.NewProgram(model_cmd{}).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
