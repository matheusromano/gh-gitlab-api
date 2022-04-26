package personalAccessToken

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

func PatRevokeExample() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := git.PersonalAccessTokens.RevokePersonalAccessToken(11311430)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response Code: %s", resp.Status)
}

func PatListExampleWithUserFilter() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListPersonalAccessTokensOptions{
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: 10},
		UserID:      gitlab.Int(11311430),
	}

	personalAccessToken, _, err := git.PersonalAccessTokens.ListPersonalAccessTokens(opt)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(personalAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)
}
