package labels

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func ExamploLabel() {
	git, err := gitlab.NewClient("glpat-qE3s3PKHsiPzPJEHzpb5")
	if err != nil {
		log.Fatal(err)
	}

	// Cria uma nova Label
	l := &gitlab.CreateLabelOptions{
		Name:  gitlab.String("Minha label"),
		Color: gitlab.String("#11FF22"),
	}

	label, _, err := git.Labels.CreateLabel("matheusrbraga/meu-projeto-via-api", l)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created label: %s\nWith color: %s\n", label.Name, label.Color)

	// Lista todoas as labels
	labels, _, err := git.Labels.ListLabels("matheusrbraga/meu-projeto-via-api", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, label := range labels {
		log.Printf("found label: %s", label.Name)
	}

}
