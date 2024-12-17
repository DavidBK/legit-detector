package detectors

import (
	"fmt"
	"time"

	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/notifications"
)

type PushTimeRule struct {
	notifer notifications.Notifier
}

func NewPushTimeRule(n notifications.Notifier) *PushTimeRule {
	return &PushTimeRule{
		notifer: n,
	}
}

func (h *PushTimeRule) GetEventTypes() []github.EventType {
	return []github.EventType{github.EventTypePush}
}

const (
	startHour = 14
	endHour   = 16
)

const pushTitle = "Push Time"

func (h *PushTimeRule) Handle(event *github.Event) {
	p := event.Payload.(*github.PushPayload)
	pushedAt := p.Repository.PushedAt
	pushDate := time.Unix(pushedAt, 0)

	if pushDate.Hour() >= startHour && pushDate.Hour() <= endHour {
		repo := p.Repository.Name
		pusher := p.Pusher.Name

		message := fmt.Sprintf("Suspicious push detected at %s by %s to %s", pushDate, pusher, repo)

		notification := notifications.Notification{
			Message:      message,
			Title:        pushTitle,
			Organization: p.Organization.Login,
			Timestamp:    pushDate,
		}

		h.notifer.Notify(notification)
	}
}
