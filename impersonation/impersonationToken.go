package impersonation

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func ExemploImpersonation() {
	git, err := gitlab.NewClient("glpat-nA6m37ysBxdAGqgDZpy")
	if err != nil {
		log.Fatal(err)
	}

	uid := 1

	//lista os token impersonalizados do seu usuario
	tokens, _, err := git.Users.GetAllImpersonationTokens(
		uid,
		&gitlab.GetAllImpersonationTokensOptions{State: gitlab.String("active")},
	)

	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		log.Println(token.Token)
	}

	// cria um token impersolizado para o usuario
	token, _, err := git.Users.CreateImpersonationToken(
		uid,
		&gitlab.CreateImpersonationTokenOptions{Scopes: &[]string{"api"}},
	)
	if err != nil {
		panic(err)
	}

	log.Println(token.Token)
}
