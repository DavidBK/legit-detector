package events

import (
	"sync"

	"github.com/davidbk6/legit-detector/detectors"
	"github.com/davidbk6/legit-detector/github"
)

type Detector = detectors.Detector

type EventDispatcher struct {
	Subscribers map[github.EventType][]Detector
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Subscribers: make(map[github.EventType][]Detector),
	}
}

func (d *EventDispatcher) Subscribe(subscriber Detector) {
	for _, eventType := range subscriber.GetEventTypes() {
		d.Subscribers[eventType] = append(d.Subscribers[eventType], subscriber)
	}
}

func (d *EventDispatcher) Dispatch(event *github.Event) error {
	if subscribers, exists := d.Subscribers[event.EventType]; exists {
		var wg sync.WaitGroup

		for _, subscriber := range subscribers {
			wg.Add(1)

			go func(s Detector) {
				defer wg.Done()
				s.Handle(event)
			}(subscriber)
		}

		wg.Wait()
	}

	// TODO: Add error handling
	return nil
}
