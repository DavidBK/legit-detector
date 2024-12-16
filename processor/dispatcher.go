package processor

import (
	"sync"

	"github.com/davidbk6/legit-detector/github"
)

type EventSubscriber interface {
	GetEventTypes() []string
	Handle(*github.Event)
}

type EventDispatcher struct {
	subscribers map[string][]EventSubscriber
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		subscribers: make(map[string][]EventSubscriber),
	}
}

func (d *EventDispatcher) Subscribe(subscriber EventSubscriber) {
	for _, eventType := range subscriber.GetEventTypes() {
		d.subscribers[eventType] = append(d.subscribers[eventType], subscriber)
	}
}

func (d *EventDispatcher) Dispatch(event *github.Event) {
	if subscribers, exists := d.subscribers[event.EventType]; exists {
		var wg sync.WaitGroup

		for _, subscriber := range subscribers {
			wg.Add(1)

			go func(s EventSubscriber) {
				defer wg.Done()
				s.Handle(event)
			}(subscriber)
		}

		wg.Wait()
	}
}
