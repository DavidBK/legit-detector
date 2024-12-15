package processor

import (
	"log"
	"time"

	"github.com/davidbk6/legit-detector/github"
)

func HandleEvent(event *github.Event) error {
	log.Printf("Processing %s event", event.EventType)
	switch event.EventType {
	case "push":
		handlePush(event)
	case "team":
		handleTeam(event)
	case "repository":
		handleRepository(event)
	default:
		log.Printf("Ignoring %s event", event.EventType)
	}

	return nil
}

func handleRepository(event *github.Event) {
	p := event.Payload.(github.RepositoryPayload)
	log.Printf("payload is %v", p)
	repoName := p.Repository.Name
	log.Printf("Processing repository event from %s", repoName)
	if p.Action == "deleted" {
		log.Printf("Repository is not legit")
	} else {
		log.Printf("Repository is legit")
	}
}

func handleTeam(event *github.Event) {
	p := event.Payload.(github.TeamPayload)
	log.Printf("payload is %v", p)
	teamName := p.Team.Name
	log.Printf("Processing team event from %s", teamName)

	// check if start with hack__
	if teamName[:5] == "hack__" {
		log.Printf("Team is not legit")
	} else {
		log.Printf("Team is legit")
	}
}

func handlePush(event *github.Event) {
	p := event.Payload.(github.PushPayload)
	log.Printf("payload is %v", p)
	pushedAt := p.Repository.PushedAt

	pushDate := time.Unix(pushedAt, 0)

	log.Printf("Processing push event from %s", pushDate)

	if pushDate.Hour() >= 14 && pushDate.Hour() <= 16 {
		log.Printf("Push is not legit")
	} else {
		log.Printf("Push is legit")
	}
}
