package processor

import (
	"log"

	"github.com/davidbk6/legit-detector/detectors"
	"github.com/davidbk6/legit-detector/github"
)

var dispatcher *EventDispatcher

func init() {
	dispatcher = NewEventDispatcher()

	// Register handlers
	dispatcher.Subscribe(detectors.NewPushHandler())
	dispatcher.Subscribe(detectors.NewTeamHandler())
	dispatcher.Subscribe(detectors.NewRepositoryHandler())
}

func HandleEvent(event *github.Event) error {
	log.Printf("Processing %s event", event.EventType)
	dispatcher.Dispatch(event)
	log.Printf("Finished processing %s event", event.EventType)
	return nil
}
