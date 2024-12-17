package detectors

import (
	"log"
	"time"

	"github.com/davidbk6/legit-detector/github"
)

type PushTimeRule struct{}

func NewPushTimeRule() *PushTimeRule {
	return &PushTimeRule{}
}

func (h *PushTimeRule) GetEventTypes() []string {
	return []string{"push"}
}

func (h *PushTimeRule) Handle(event *github.Event) {
	p := event.Payload.(*github.PushPayload)
	pushedAt := p.Repository.PushedAt
	pushDate := time.Unix(pushedAt, 0)

	log.Printf("Processing push event from %s", pushDate)
	log.Printf("Push organization: %s", p.Organization.Login)
	if pushDate.Hour() >= 14 && pushDate.Hour() <= 16 {
		log.Printf("Push is not legit")
	} else {
		log.Printf("Push is legit")
	}
}
