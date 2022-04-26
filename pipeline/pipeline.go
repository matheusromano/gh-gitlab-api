package pipeline

import (
	"log"
	"time"

	"github.com/xanzy/go-gitlab"
)

func Pipeline() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectPipelinesOptions{
		Scope:         gitlab.String("branches"),
		Status:        gitlab.BuildState(gitlab.Created),
		Ref:           gitlab.String("main"),
		YamlErrors:    gitlab.Bool(true),
		Name:          gitlab.String("TestGoPipeline"),
		Username:      gitlab.String("matheusrbraga"),
		UpdatedAfter:  gitlab.Time(time.Now().Add(-24 * 365 * time.Hour)),
		UpdatedBefore: gitlab.Time(time.Now().Add(-7 * 24 * time.Hour)),
		OrderBy:       gitlab.String("status"),
		Sort:          gitlab.String("asc"),
	}

	pipelines, _, err := git.Pipelines.ListProjectPipelines(35180345, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines {
		log.Printf("found pipeline: %v", pipeline)
	}
}
