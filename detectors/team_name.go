package detectors

import (
	"fmt"
	"strings"
	"time"

	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/notifications"
)

type TeamNameRule struct{}

func NewTeamNameRule() *TeamNameRule {
	return &TeamNameRule{}
}

func (h *TeamNameRule) GetEventTypes() []string {
	return []string{"team"}
}

func (h *TeamNameRule) Handle(event *github.Event) {
	p := event.Payload.(*github.TeamPayload)
	teamName := p.Team.Name

	if strings.HasPrefix(teamName, "hacker") {
		message := fmt.Sprintf("Suspicious team name detected: %s, %s by %s", teamName, p.Action, p.Sender.Login)

		notification := notifications.Notification{
			Message:      message,
			EventType:    "team",
			Organization: p.Organization.Login,
			Timestamp:    time.Now(),
		}

		notifications.GetManager().NotifyAll(notification)
	}
}
