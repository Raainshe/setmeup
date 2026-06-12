package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/raainshe/setmeup/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialiseModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
