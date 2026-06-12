package tui

import (
	"os"
	"os/exec"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/raainshe/setmeup/templates"
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
	if err := exec.Command("npm", args...).Run(); err != nil {
		return err
	}

	port := m.FrontendPort
	if port == "" {
		port = "5173"
	}

	frontendDir := m.FrontendName
	if frontendDir == "" {
		frontendDir = "frontend"
	}

	dockerfile := templates.Render(templates.FrontendDocker, map[string]string{
		"FRONTEND_PORT": port,
	})
	if err := os.WriteFile(filepath.Join(frontendDir, "Dockerfile"), []byte(dockerfile), 0o644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(frontendDir, ".dockerignore"), []byte(templates.FrontendDockerignore), 0o644); err != nil {
		return err
	}

	return nil
}

func (m Model) CreateBackend() error {
	args := []string{
		"create",
		"--name", m.BackendName,
		"--framework", "gin",
		"--driver", "mongo",
		"--feature", "docker",
		"--git", "skip",
	}
	err := exec.Command("go-blueprint", args...).Run()
	if err != nil {
		return err
	}

	return nil
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
