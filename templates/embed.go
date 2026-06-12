package templates

import (
	_ "embed"
	"strings"
)

//go:embed frontend/docker.tmpl
var FrontendDocker string

//go:embed backend/docker.tmpl
var BackendDocker string

//go:embed compose.yml.tmpl
var ComposeYML string

//go:embed Makefile.tmpl
var Makefile string

const FrontendDockerignore = `node_modules
dist
npm-debug.log
.DS_Store
.git
.gitignore
.env
*.env
`

const BackendDockerignore = `
.env
*.env
data/
uploads/
*.log
dist/
.git
.gitignore
tmp/
`

func Render(tmpl string, vars map[string]string) string {
	out := tmpl
	for key, val := range vars {
		out = strings.ReplaceAll(out, "{{"+key+"}}", val)
	}
	return out
}
