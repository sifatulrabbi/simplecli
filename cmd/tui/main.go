package main

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/sifatulrabbi/simplecli/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.NewModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
