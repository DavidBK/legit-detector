package detectors

import (
	"log"

	"github.com/davidbk6/legit-detector/github"
)

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

func (h *Ping) GetEventTypes() []string {
	return []string{"ping"}
}

func (h *Ping) Handle(event *github.Event) {
	log.Printf("Processing ping event")
	log.Printf("New Webhook %s", event.Payload)
}
