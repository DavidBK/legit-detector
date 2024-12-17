package detectors

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/notifications"
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

				repo := p.Repository.Name
				sender := p.Sender.Login
				message := fmt.Sprintf("Suspiciously short lifetime of %s detected (%s), sender: %s", repo, lifetime, sender)

				notification := notifications.Notification{
					Message:      message,
					EventType:    "repository",
					Organization: p.Organization.Login,
					Timestamp:    deleteTime,
				}

				notifications.GetManager().NotifyAll(notification)
			}
		} else {
			log.Printf("Repository %s deletion detected but creation time unknown", repoName)
		}

	default:
		log.Printf("Repository %s: other action (%s) detected", repoName, p.Action)
	}
}
