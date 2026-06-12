package tui

import (
	"os"
	"os/exec"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/raainshe/setmeup/templates"
)

func writeRendered(path, tmpl string, vars map[string]string) error {
	return os.WriteFile(path, []byte(templates.Render(tmpl, vars)), 0o644)
}

func (m Model) CreateFrontEnd() error {
	vars := templateVars(m)

	args := []string{
		"create",
		"vue@latest",
		"--",
		vars["FRONTEND_NAME"],
		"--ts",
		"--router",
		"--pinia",
		"--eslint",
		"--prettier",
	}
	if err := exec.Command("npm", args...).Run(); err != nil {
		return err
	}

	frontendDir := vars["FRONTEND_NAME"]

	if err := writeRendered(
		filepath.Join(frontendDir, "Dockerfile"),
		templates.FrontendDocker,
		map[string]string{"FRONTEND_PORT": vars["FRONTEND_PORT"]},
	); err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Join(frontendDir, ".dockerignore"),
		[]byte(templates.FrontendDockerignore),
		0o644,
	)
}

func (m Model) injectBackend(vars map[string]string) error {
	backendDir := vars["BACKEND_NAME"]

	if err := writeRendered(
		filepath.Join(backendDir, "Dockerfile"),
		templates.BackendDocker,
		map[string]string{"BACKEND_PORT": vars["BACKEND_PORT"]},
	); err != nil {
		return err
	}

	if err := os.WriteFile(
		filepath.Join(backendDir, ".dockerignore"),
		[]byte(templates.BackendDockerignore),
		0o644,
	); err != nil {
		return err
	}

	if err := writeRendered(filepath.Join(backendDir, ".env.example"), templates.BackendEnvExample, vars); err != nil {
		return err
	}

	if err := writeRendered(filepath.Join(backendDir, ".env"), templates.BackendEnv, vars); err != nil {
		return err
	}

	if err := writeRendered(
		filepath.Join(backendDir, "internal", "database", "database.go"),
		templates.BackendDatabaseGo,
		vars,
	); err != nil {
		return err
	}

	if err := writeRendered(
		filepath.Join(backendDir, "internal", "database", "database_test.go"),
		templates.BackendDatabaseTestGo,
		vars,
	); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(backendDir, "docker-compose.yml")); err != nil && !os.IsNotExist(err) {
		return err
	}

	makefilePath := filepath.Join(backendDir, "Makefile")
	makefile, err := os.ReadFile(makefilePath)
	if err != nil {
		return err
	}
	if err := os.WriteFile(makefilePath, []byte(patchMakefile(string(makefile))), 0o644); err != nil {
		return err
	}

	readmePath := filepath.Join(backendDir, "README.md")
	readme, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}
	if err := os.WriteFile(readmePath, []byte(patchREADME(string(readme))), 0o644); err != nil {
		return err
	}

	return nil
}

func (m Model) CreateBackend() error {
	vars := templateVars(m)

	args := []string{
		"create",
		"--name", vars["BACKEND_NAME"],
		"--framework", "gin",
		"--driver", "mongo",
		"--git", "skip",
	}
	if err := exec.Command("go-blueprint", args...).Run(); err != nil {
		return err
	}

	return m.injectBackend(vars)
}

func (m Model) CreateWorkspace() tea.Msg {
	if err := m.CreateFrontEnd(); err != nil {
		return errMsg(err)
	}
	if err := m.CreateBackend(); err != nil {
		return errMsg(err)
	}

	vars := templateVars(m)

	if err := writeRendered("docker-compose.yml", templates.ComposeYML, vars); err != nil {
		return errMsg(err)
	}

	if err := writeRendered("Makefile", templates.Makefile, vars); err != nil {
		return errMsg(err)
	}

	return generateMsg("Workspace created successfully")
}
