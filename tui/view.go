package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

// Define reusable UI styles using Lip Gloss
var (
	// Vibrant purple banner style for the application header
	titleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#6200EE")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2).
			Bold(true).
			MarginBottom(1)

	// Bright green style for successful operations
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true)

	// Crimson style for errors
	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true)

	// Highlighting style for active interactive elements
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true)

	// Subtle gray formatting style for help context details
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)
)

// View renders the terminal screen based on the current step state
func (m Model) View() tea.View {
	// 1. Build the permanent application header banner
	var b strings.Builder
	b.WriteString(titleStyle.Render(" SETMEUP: FULL-STACK SCAFFOLDER ") + "\n")

	// 2. Render content blocks conditionally based on the active step
	switch m.Step {
	case StepWelcome:
		b.WriteString("Do you want to customize folder names and port numbers?\n\n")

		yesOption := "  [ ] Yes, let me change things"
		noOption := "  [ ] No, use fallback defaults (frontend:5173, backend:8080)"

		// Visually inject checkmarks and color variations based on cursor focus
		if m.CursorChoice == 0 {
			yesOption = selectedStyle.Render("  [x] Yes, let me change things")
		} else {
			noOption = selectedStyle.Render("  [x] No, use fallback defaults (frontend:5173, backend:8080)")
		}

		b.WriteString(yesOption + "\n" + noOption + "\n\n")
		b.WriteString(helpStyle.Render("(Use Arrow Keys or H/J/K/L to toggle, press Enter to select)"))

	case StepBackendName:
		b.WriteString("📁 Backend Configuration\n\n")
		b.WriteString(fmt.Sprintf("Enter Backend Directory Name (default: backend):\n%s", m.Inputs[0].View()))

	case StepBackendPort:
		b.WriteString("🔌 Backend Configuration\n\n")
		b.WriteString(fmt.Sprintf("Enter Backend Development Port (default: 8080):\n%s", m.Inputs[1].View()))

	case StepFrontendName:
		b.WriteString("📁 Frontend Configuration\n\n")
		b.WriteString(fmt.Sprintf("Enter Frontend Directory Name (default: frontend):\n%s", m.Inputs[2].View()))

	case StepFrontendPort:
		b.WriteString("🔌 Frontend Configuration\n\n")
		b.WriteString(fmt.Sprintf("Enter Frontend Development Port (default: 5173):\n%s", m.Inputs[3].View()))

	case StepGenerating:
		b.WriteString("⏳ Orchestrating project directories...\n\n")
		b.WriteString(fmt.Sprintf("  • UI Folder:   ./%s (Port %s)\n", m.FrontendName, m.FrontendPort))
		b.WriteString(fmt.Sprintf("  • API Folder:  ./%s (Port %s)\n\n", m.BackendName, m.BackendPort))
		b.WriteString(helpStyle.Render("Running go-blueprint templates and executing injections..."))

	case StepDone:
		if m.Err != nil {
			b.WriteString(errorStyle.Render("Generation Failed!") + "\n")
			b.WriteString(fmt.Sprintf("Reason: %v\n", m.Err))
		} else {
			b.WriteString(successStyle.Render("Workspace successfully created!") + "\n\n")
			b.WriteString(m.Logs + "\n")
		}
	}

	// 3. Append global footer exit instructions
	b.WriteString("\n\n" + helpStyle.Render("[press ctrl+c or q to quit]") + "\n")

	return tea.NewView(b.String())
}
