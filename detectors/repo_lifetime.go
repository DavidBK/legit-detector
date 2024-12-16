package detectors

import (
	"log"

	"github.com/davidbk6/legit-detector/github"
)

type RepoLifeTimeRule struct{}

func NewRepoLifeTimeRule() *RepoLifeTimeRule {
	return &RepoLifeTimeRule{}
}

func (h *RepoLifeTimeRule) GetEventTypes() []string {
	return []string{"repository"}
}

func (h *RepoLifeTimeRule) Handle(event *github.Event) {
	p := event.Payload.(*github.RepositoryPayload)
	repoName := p.Repository.Name

	log.Printf("Processing repository event from %s", repoName)
	if p.Action == "deleted" {
		log.Printf("Repository is not legit")
	} else {
		log.Printf("Repository is legit")
	}
}
