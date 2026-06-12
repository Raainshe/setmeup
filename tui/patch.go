package tui

import "regexp"

var (
	makefileDockerRunRe  = regexp.MustCompile(`(?m)^docker-run:\n(?:\t.*\n)+`)
	makefileDockerDownRe = regexp.MustCompile(`(?m)^docker-down:\n(?:\t.*\n)+`)
	readmeDockerSectionRe = regexp.MustCompile(`(?s)Create DB container\n` + "```bash\nmake docker-run\n```\n\n" + `Shutdown DB Container\n` + "```bash\nmake docker-down\n```\n\n")
)

const readmeDockerReplacement = `## Docker (root compose)

From the workspace root, start backend and frontend:

` + "```bash\nmake up\n```\n\n" + `Stop containers:

` + "```bash\nmake down\n```\n\n" + `Configure MongoDB Atlas in backend/.env — copy backend/.env.example and set MONGO_URI before starting.

`

func patchMakefile(content string) string {
	content = makefileDockerRunRe.ReplaceAllString(content, "docker-run:\n\t$(MAKE) -C .. up\n\n")
	content = makefileDockerDownRe.ReplaceAllString(content, "docker-down:\n\t$(MAKE) -C .. down\n\n")
	return content
}

func patchREADME(content string) string {
	if readmeDockerSectionRe.MatchString(content) {
		return readmeDockerSectionRe.ReplaceAllString(content, readmeDockerReplacement)
	}
	return content
}
