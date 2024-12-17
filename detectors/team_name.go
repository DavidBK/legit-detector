package detectors

import (
	"log"
	"strings"

	"github.com/davidbk6/legit-detector/github"
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

	log.Printf("Processing team event from %s", teamName)
	log.Printf("Organization: %s", p.Organization.Login)

	if strings.HasPrefix(teamName, "hacker") {
		log.Printf("Team is not legit")
	} else {
		log.Printf("Team is legit")
	}
}
