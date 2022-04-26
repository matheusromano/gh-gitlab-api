package application

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func CreateApplication() {
	git, err := gitlab.NewClient("glpat-nA6m37ysBxdAGqgDZpyJ")
	if err != nil {
		log.Fatal(err)
	}

	// Create an Application
	opts := &gitlab.CreateApplicationOptions{
		Name:        gitlab.String("Travis"),
		RedirectURI: gitlab.String("http://example.org"),
		Scopes:      gitlab.String("api"),
	}

	created, _, err := git.Applications.CreateApplication(opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Ultima aplicação criada: %v", created)

	applications, _, err := git.Applications.ListApplications(&gitlab.ListApplicationsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, app := range applications {
		log.Printf("Found app: %v", app)
	}

	// Delete uma aplicação
	resp, err := git.Applications.DeleteApplication(created.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Status code response : %d", resp.StatusCode)
}
