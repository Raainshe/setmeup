package tui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type generateMsg string
type errMsg error

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Step == StepWelcome {
				m.CursorChoice = 0 // Yes — custom wizard
			}

		case "down", "j":
			if m.Step == StepWelcome {
				m.CursorChoice = 1 // No — use defaults
			}

		case "enter":
			if m.Step == StepWelcome {
				if m.CursorChoice == 1 {
					m.FrontendName = "frontend"
					m.FrontendPort = "5173"
					m.BackendName = "backend"
					m.BackendPort = "8080"
					m.Step = StepGenerating
					return m, m.CreateWorkspace
				}
				m.Step = StepBackendName
				return m, tea.Batch(textinput.Blink, m.Inputs[0].Focus())
			}

			if m.Step < StepFrontendPort {
				m.Inputs[m.Step-1].Blur()
				m.Step++
				return m, tea.Batch(textinput.Blink, m.Inputs[m.Step-1].Focus())
			} else if m.Step == StepFrontendPort {
				m.BackendName = m.Inputs[0].Value()
				m.BackendPort = m.Inputs[1].Value()
				m.FrontendName = m.Inputs[2].Value()
				m.FrontendPort = m.Inputs[3].Value()

				//if user left an empty string
				if m.BackendName == "" {
					m.BackendName = "backend"
				}
				if m.BackendPort == "" {
					m.BackendPort = "8080"
				}
				if m.FrontendName == "" {
					m.FrontendName = "frontend"
				}
				if m.FrontendPort == "" {
					m.FrontendPort = "5173"
				}
				m.Step = StepGenerating
				return m, tea.Printf("Generating Code")
			}
		}
	case generateMsg:
		m.Step = StepDone
		m.Logs = string(msg)
		return m, tea.Quit
	case errMsg:
		m.Err = msg
		m.Step = StepDone
		return m, nil
	}

	if m.Step >= StepBackendName && m.Step <= StepFrontendPort {
		var cmd tea.Cmd
		m.Inputs[m.Step-1], cmd = m.Inputs[m.Step-1].Update(msg)
		return m, cmd
	}
	return m, nil
}
