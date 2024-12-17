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
	notifer           notifications.Notifier
}

func NewRepoLifeTimeRule(n notifications.Notifier) *RepoLifeTimeRule {
	return &RepoLifeTimeRule{
		repoCreationTimes: make(map[int]time.Time),
		minimumLifetime:   lifetimeThreshold,
		notifer:           n,
	}
}

func (r *RepoLifeTimeRule) GetEventTypes() []github.EventType {
	return []github.EventType{github.EventTypeRepository}
}

func (r *RepoLifeTimeRule) Handle(event *github.Event) {
	p := event.Payload.(*github.RepositoryPayload)
	repoID := p.Repository.ID
	repoName := p.Repository.Name
	creationTime := p.Repository.CreatedAt

	switch p.Action {
	case "created":
		r.mu.Lock()
		r.repoCreationTimes[repoID] = creationTime
		r.mu.Unlock()
		log.Printf("Repository %s created and tracked", repoName)

	case "deleted":
		r.mu.Lock()
		creationTime, exists := r.repoCreationTimes[repoID]
		if exists {
			delete(r.repoCreationTimes, repoID)
		}
		r.mu.Unlock()

		if exists {
			deleteTime := p.Repository.UpdatedAt
			lifetime := deleteTime.Sub(creationTime)

			if lifetime < r.minimumLifetime {

				repo := p.Repository.Name
				sender := p.Sender.Login
				message := fmt.Sprintf("Suspiciously short lifetime of %s detected (%s), sender: %s", repo, lifetime, sender)

				notification := notifications.Notification{
					Message:      message,
					EventType:    "repository",
					Organization: p.Organization.Login,
					Timestamp:    deleteTime,
				}

				r.notifer.Notify(notification)
			}
		} else {
			log.Printf("Repository %s deletion detected but creation time unknown", repoName)
		}

	default:
		log.Printf("Repository %s: other action (%s) detected", repoName, p.Action)
	}
}
