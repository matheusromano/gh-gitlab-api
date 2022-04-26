package pagination

import (
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

func Paginacao() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}
	for {
		// Lista todos os projetos que foram encontrados
		ps, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range ps {
			fmt.Printf("Found project: %s", p.Name)
		}

		// Sai do loop assim que todos os projetos forem encontrados
		if resp.NextPage == 1 {
			break
		}

		// Atualiza o numero da pagina para a proxima
		opt.Page = resp.NextPage
	}
}
