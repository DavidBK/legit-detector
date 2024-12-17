package detectors

import (
	"log"
	"sync"
	"time"

	"github.com/davidbk6/legit-detector/github"
)

const lifetimeThreshold = 10 * time.Minute

// TODO: Add cleanup for old repo creation times

type RepoLifeTimeRule struct {
	mu                sync.RWMutex
	repoCreationTimes map[int]time.Time
	minimumLifetime   time.Duration
}

func NewRepoLifeTimeRule() *RepoLifeTimeRule {
	return &RepoLifeTimeRule{
		repoCreationTimes: make(map[int]time.Time),
		minimumLifetime:   lifetimeThreshold,
	}
}

func (h *RepoLifeTimeRule) GetEventTypes() []string {
	return []string{"repository"}
}

func (h *RepoLifeTimeRule) Handle(event *github.Event) {
	p := event.Payload.(*github.RepositoryPayload)
	repoID := p.Repository.ID
	repoName := p.Repository.Name
	creationTime := p.Repository.CreatedAt

	log.Printf("Processing repository event from %s (ID: %d)", repoName, repoID)
	log.Printf("Organization: %s", p.Organization.Login)

	switch p.Action {
	case "created":
		h.mu.Lock()
		h.repoCreationTimes[repoID] = creationTime
		h.mu.Unlock()
		log.Printf("Repository %s created and tracked", repoName)

	case "deleted":
		h.mu.Lock()
		creationTime, exists := h.repoCreationTimes[repoID]
		if exists {
			delete(h.repoCreationTimes, repoID)
		}
		h.mu.Unlock()

		if exists {
			deleteTime := p.Repository.UpdatedAt
			lifetime := deleteTime.Sub(creationTime)

			if lifetime < h.minimumLifetime {
				log.Printf("Repository %s is not legit: lived for only %v (less than %v)",
					repoName, lifetime.Round(time.Second), h.minimumLifetime)

			} else {
				log.Printf("Repository %s is legit: lived for %v",
					repoName, lifetime.Round(time.Second))
			}
		} else {
			log.Printf("Repository %s deletion detected but creation time unknown", repoName)
		}

	default:
		log.Printf("Repository %s: other action (%s) detected", repoName, p.Action)
	}
}
