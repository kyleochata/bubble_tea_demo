package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model_b struct {
	choices  []string
	cursor   int
	selected map[int]struct{} //holds what choices are selected
}

func initModel() model_b {
	return model_b{
		choices:  []string{"Buy lettuce", "Buy celery", "Buy miso"},
		selected: make(map[int]struct{}),
	}
}
func (m model_b) Init() tea.Cmd {
	return nil
}
func (m model_b) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}
func (m model_b) View() string {
	//header
	s := "What should we buy at the market?\n\n"
	for i, choice := range m.choices {
		cursor := " " //no cursor
		//is cursor pointing at this choice?
		if m.cursor == i {
			cursor = ">" //cursor!
		}
		//is choice selected?
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x" //selected!
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress ctrl+c or q to quit. \n"
	return s
}

func main1() {
	program := tea.NewProgram(initModel())
	if _, err := program.Run(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
