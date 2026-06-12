package templates

import (
	"strings"
	"testing"
)

func TestRenderBackendTemplates(t *testing.T) {
	vars := map[string]string{
		"BACKEND_NAME":  "backend",
		"BACKEND_PORT":  "8080",
		"FRONTEND_NAME": "frontend",
		"FRONTEND_PORT": "5173",
	}

	cases := map[string]string{
		"BackendDocker":         BackendDocker,
		"BackendEnvExample":     BackendEnvExample,
		"BackendEnv":            BackendEnv,
		"BackendDatabaseGo":     BackendDatabaseGo,
		"BackendDatabaseTestGo": BackendDatabaseTestGo,
		"ComposeYML":            ComposeYML,
		"FrontendDocker":        FrontendDocker,
	}

	for name, tmpl := range cases {
		out := Render(tmpl, vars)
		if out == "" {
			t.Fatalf("%s rendered empty", name)
		}
		if containsPlaceholder(out) {
			t.Fatalf("%s still contains placeholders:\n%s", name, out)
		}
	}
}

func containsPlaceholder(s string) bool {
	return strings.Contains(s, "{{BACKEND") || strings.Contains(s, "{{FRONTEND")
}
