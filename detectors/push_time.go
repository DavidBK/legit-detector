package detectors

import (
	"log"
	"time"

	"github.com/davidbk6/legit-detector/github"
)

type PushHandler struct{}

func NewPushHandler() *PushHandler {
	return &PushHandler{}
}

func (h *PushHandler) GetEventTypes() []string {
	return []string{"push"}
}

func (h *PushHandler) Handle(event *github.Event) {
	p := event.Payload.(*github.PushPayload)
	pushedAt := p.Repository.PushedAt
	pushDate := time.Unix(pushedAt, 0)

	log.Printf("Processing push event from %s", pushDate)
	if pushDate.Hour() >= 14 && pushDate.Hour() <= 16 {
		log.Printf("Push is not legit")
	} else {
		log.Printf("Push is legit")
	}
}
