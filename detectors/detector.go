package detectors

import "github.com/davidbk6/legit-detector/github"

type Detector interface {
	GetEventTypes() []github.EventType
	Handle(*github.Event)
}
