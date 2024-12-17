package registry

import (
	"github.com/davidbk6/legit-detector/github"
)

var defaultRegistry = NewRegistry()

type EventSubscriber interface {
	GetEventTypes() []string
	Handle(*github.Event)
}

type Registry struct {
	subscribers []EventSubscriber
}

func NewRegistry() *Registry {
	return &Registry{
		subscribers: make([]EventSubscriber, 0),
	}
}

func (r *Registry) Register(subscriber EventSubscriber) {
	r.subscribers = append(r.subscribers, subscriber)
}

func (r *Registry) GetSubscribers() []EventSubscriber {
	return r.subscribers
}
