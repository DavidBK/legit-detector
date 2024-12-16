package detectors

import (
	"log"

	"github.com/davidbk6/legit-detector/github"
)

type RepositoryHandler struct{}

func NewRepositoryHandler() *RepositoryHandler {
	return &RepositoryHandler{}
}

func (h *RepositoryHandler) GetEventTypes() []string {
	return []string{"repository"}
}

func (h *RepositoryHandler) Handle(event *github.Event) {
	p := event.Payload.(*github.RepositoryPayload)
	repoName := p.Repository.Name

	log.Printf("Processing repository event from %s", repoName)
	if p.Action == "deleted" {
		log.Printf("Repository is not legit")
	} else {
		log.Printf("Repository is legit")
	}
}
