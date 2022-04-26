package createproject

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func CreateProject() {

	git, err := gitlab.NewClient("glpat-wkKKuJ6f4ZPz7A8xpHzR")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Criando um projeto
	p := &gitlab.CreateProjectOptions{
		Name:                 gitlab.String("Meu projeto via API"),
		Description:          gitlab.String("Apenas um projeto para testar a API"),
		MergeRequestsEnabled: gitlab.Bool(true),
		SnippetsEnabled:      gitlab.Bool(true),
		Visibility:           gitlab.Visibility(gitlab.PublicVisibility),
	}

	projects, _, err := git.Projects.CreateProject(p)
	if err != nil {
		log.Fatal(err)
	}

	// Adcionando um novo snippet
	s := &gitlab.CreateProjectSnippetOptions{
		Title:      gitlab.String("Snippet Bobo"),
		FileName:   gitlab.String("snippet.go"),
		Content:    gitlab.String("package main...."),
		Visibility: gitlab.Visibility(gitlab.PublicVisibility),
	}

	_, _, err = git.ProjectSnippets.CreateSnippet(projects.ID, s)
	if err != nil {
		log.Fatal(err)
	}

}
