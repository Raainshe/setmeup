package tui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type Step int

const (
	StepWelcome Step = iota
	StepBackendName
	StepBackendPort
	StepFrontendName
	StepFrontendPort
	StepGenerating
	StepDone
)

type Model struct {
	Step         Step
	Inputs       []textinput.Model
	BackendName  string
	BackendPort  string
	FrontendName string
	FrontendPort string
	CustomConfig bool
	CursorChoice int
	Logs         string
	Err          error
}

func InitialiseModel() Model {

	m := Model{
		Step:         StepWelcome,
		Inputs:       make([]textinput.Model, 4),
		CursorChoice: 0,
	}

	//backend name
	m.Inputs[0] = textinput.New()
	m.Inputs[0].Placeholder = "backend"

	//backend port
	m.Inputs[1] = textinput.New()
	m.Inputs[1].Placeholder = "8080"

	//backend name
	m.Inputs[2] = textinput.New()
	m.Inputs[2].Placeholder = "frontend"

	//backend name
	m.Inputs[3] = textinput.New()
	m.Inputs[3].Placeholder = "5173"

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
