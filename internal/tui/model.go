// Package tui for handling the terminal models
package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Model defines the state of our TUI
// You can add more fields as needed
// This example is a simple menu

// MenuItem represents a selectable item in the menu
// Add more fields for more advanced functionality

type MenuItem struct {
	Title string
	Desc  string
	Exec  func() error
}

type Model struct {
	menu        []MenuItem
	selectedIdx int
	quitting    bool
	textInput   textinput.Model
}

func NewModel() Model {
	return Model{
		menu: []MenuItem{
			{
				Title: "Minify",
				Desc:  "Minify HTML, CSS, or JS file",
			},
			{
				Title: "Snippets",
				Desc:  "Generate a framework snippet",
			},
			{
				Title: "Quit",
				Desc:  "Exit the application",
			},
		},
		selectedIdx: 0,
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.selectedIdx > 0 {
				m.selectedIdx--
			}
		case "down", "j":
			if m.selectedIdx < len(m.menu)-1 {
				m.selectedIdx++
			}
		case "enter":
			if m.menu[m.selectedIdx].Title == "Quit" {
				m.quitting = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}
	out := ""
	for i, item := range m.menu {
		cursor := "[ ]"
		if m.selectedIdx == i {
			cursor = "[x]"
		}
		out += cursor + " " + item.Title + " - " + item.Desc + "\n"
	}
	out += "\n↑/k and ↓/j to move, Enter to select, q to quit\n"
	return out
}
