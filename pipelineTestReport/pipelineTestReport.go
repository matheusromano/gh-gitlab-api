package pipelineTestReport

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func PipelineTestReport() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectPipelinesOptions{Ref: gitlab.String("main")}
	projectID := 35180345

	pipelines, _, err := git.Pipelines.ListProjectPipelines(projectID, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines {
		log.Printf("Found pipeline: %v", pipeline)

		report, _, err := git.Pipelines.GetPipelineTestReport(projectID, pipeline.ID)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found test resport: %v", report)
	}

}
