package detectors

import (
	"log"

	"github.com/davidbk6/legit-detector/github"
)

type TeamHandler struct{}

func NewTeamHandler() *TeamHandler {
	return &TeamHandler{}
}

func (h *TeamHandler) GetEventTypes() []string {
	return []string{"team"}
}

func (h *TeamHandler) Handle(event *github.Event) {
	p := event.Payload.(*github.TeamPayload)
	teamName := p.Team.Name

	log.Printf("Processing team event from %s", teamName)

	if len(teamName) >= 5 && teamName[:5] != "hack__" {
		log.Printf("Team is not legit")
	} else {
		log.Printf("Team is legit")
	}
}
