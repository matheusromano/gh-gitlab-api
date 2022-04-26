package repositoryFiles

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func RepositoryFiles() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	cf := &gitlab.CreateFileOptions{
		Branch:        gitlab.String("main"),
		Content:       gitlab.String("my file contents"),
		CommitMessage: gitlab.String("feat: add new content via api"),
	}

	file, _, err := git.RepositoryFiles.CreateFile("matheusrbraga/meu-projeto-via-api", "file.go", cf)
	if err != nil {
		log.Fatal(err)
	}

	// Atualiza o arquivos do repositorio
	uf := &gitlab.UpdateFileOptions{
		Branch:        gitlab.String("main"),
		Content:       gitlab.String("Test Content"),
		CommitMessage: gitlab.String("fix: something was wrong"),
	}

	_, _, err = git.RepositoryFiles.UpdateFile("matheusrbraga/meu-projeto-via-api", file.FilePath, uf)
	if err != nil {
		log.Fatal(err)
	}

	gf := &gitlab.GetFileOptions{
		Ref: gitlab.String("main"),
	}

	f, _, err := git.RepositoryFiles.GetFile("matheusrbraga/meu-projeto-via-api", file.FilePath, gf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contains: %s", f.Content)

	gfb := &gitlab.GetFileBlameOptions{
		Ref: gitlab.String("main"),
	}

	fb, _, err := git.RepositoryFiles.GetFileBlame("matheusrbraga/meu-projeto-via-api", file.FilePath, gfb)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d blame ranges", len(fb))

}
