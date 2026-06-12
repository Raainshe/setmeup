package tui

import (
	"strings"
	"testing"
)

func TestPatchMakefile(t *testing.T) {
	input := `# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
`
	got := patchMakefile(input)
	for _, want := range []string{"docker-run:\n\t$(MAKE) -C .. up", "docker-down:\n\t$(MAKE) -C .. down"} {
		if !strings.Contains(got, want) {
			t.Fatalf("patchMakefile missing %q:\n%s", want, got)
		}
	}
}

func TestPatchREADME(t *testing.T) {
	input := `Run the application
` + "```bash\nmake run\n```\n\n" + `Create DB container
` + "```bash\nmake docker-run\n```\n\n" + `Shutdown DB Container
` + "```bash\nmake docker-down\n```\n\n" + `DB Integrations Test:
` + "```bash\nmake itest\n```\n"

	got := patchREADME(input)
	for _, want := range []string{"make up", "make down", "MONGO_URI", "backend/.env.example"} {
		if !strings.Contains(got, want) {
			t.Fatalf("patchREADME missing %q:\n%s", want, got)
		}
	}
}
