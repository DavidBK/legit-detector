package processor

import (
	"log"

	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/registry"
)

type Processor struct {
	dispatcher *EventDispatcher
}

func NewProcessor(registry *registry.Registry) *Processor {
	dispatcher := NewEventDispatcher()

	for _, subscriber := range registry.GetSubscribers() {
		dispatcher.Subscribe(subscriber)
	}

	return &Processor{
		dispatcher: dispatcher,
	}
}

func (p *Processor) HandleEvent(event *github.Event) error {
	log.Printf("Processing %s event", event.EventType)
	err := p.dispatcher.Dispatch(event)
	if err != nil {
		return err
	}
	log.Printf("Finished processing %s event", event.EventType)
	return nil
}
