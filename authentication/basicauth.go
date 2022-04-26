package authentication

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func ExemploBasicoAutenticacao() {

	git, err := gitlab.NewBasicAuthClient(
		"matheusrbraga",
		"MAthbzica93",
		gitlab.WithBaseURL("https://gitlab.com"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d projects", len(projects))
}
