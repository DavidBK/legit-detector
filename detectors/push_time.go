package detectors

import (
	"fmt"
	"time"

	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/notifications"
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

	if pushDate.Hour() >= 14 && pushDate.Hour() <= 16 {
		repo := p.Repository.Name
		pusher := p.Pusher.Name

		message := fmt.Sprintf("Suspicious push detected at %s by %s to %s", pushDate, pusher, repo)

		notification := notifications.Notification{
			Message:      message,
			EventType:    "push",
			Organization: p.Organization.Login,
			Timestamp:    pushDate,
		}

		notifications.GetManager().NotifyAll(notification)
	}
}
