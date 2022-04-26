package languages

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func ExampleLanguage() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	languages, _, err := git.Projects.GetProjectLanguages("matheusrbraga/meu-projeto-via-api")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("found languages: %v", languages)
}
