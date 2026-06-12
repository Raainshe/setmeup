package tui

func templateVars(m Model) map[string]string {
	backendName := m.BackendName
	if backendName == "" {
		backendName = "backend"
	}

	backendPort := m.BackendPort
	if backendPort == "" {
		backendPort = "8080"
	}

	frontendName := m.FrontendName
	if frontendName == "" {
		frontendName = "frontend"
	}

	frontendPort := m.FrontendPort
	if frontendPort == "" {
		frontendPort = "5173"
	}

	return map[string]string{
		"BACKEND_NAME":  backendName,
		"BACKEND_PORT":  backendPort,
		"FRONTEND_NAME": frontendName,
		"FRONTEND_PORT": frontendPort,
	}
}
