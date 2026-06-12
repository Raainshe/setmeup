package tui

import (
	"os/exec"

	tea "charm.land/bubbletea/v2"
)

func (m Model) CreateFrontEnd() error {

	args := []string{
		"create",
		"vue@latest",
		"--",
		m.FrontendName,
		"--ts",
		"--router",
		"--pinia",
		"--eslint",
		"--prettier",
	}
	err := exec.Command("npm", args...).Run()
	return err
}

func (m Model) CreateBackend() error {
	args := []string{
		"create",
		"--name", m.BackendName, // Dynamically injects your backend folder name
		"--framework", "gin",
		"--driver", "mongo",
		"--feature", "docker",
		"--git", "skip",
	}
	err := exec.Command("go-blueprint", args...).Run()
	return err
}

func (m Model) CreateWorkspace() tea.Msg {
	err := m.CreateFrontEnd()
	if err != nil {
		return errMsg(err)
	}
	err = m.CreateBackend()
	if err != nil {
		return errMsg(err)
	}
	return generateMsg("Workspace created successfully")
}
