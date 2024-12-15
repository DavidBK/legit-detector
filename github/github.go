package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Event struct {
	EventType string `json:"event_type"`
	Payload   any    `json:"payload"`
}

type PushPayload struct {
	Repository struct {
		PushedAt int64 `json:"pushed_at"`
	} `json:"repository"`
}

type TeamPayload struct {
	Action string `json:"action"`
	Team   struct {
		Name string `json:"name"`
	} `json:"team"`
}

type RepositoryPayload struct {
	Action     string `json:"action"`
	Repository struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"repository"`
}

func ParseEvent(r *http.Request) (*Event, error) {
	eventType := r.Header.Get("X-GitHub-Event")
	if eventType == "" {
		return nil, fmt.Errorf("missing X-GitHub-Event header")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	switch eventType {
	case "push":
		var payload PushPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %v", err)
		}

		return &Event{
			EventType: eventType,
			Payload:   payload,
		}, nil

	case "team":
		var payload TeamPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %v", err)
		}
		return &Event{
			EventType: eventType,
			Payload:   payload,
		}, nil

	case "repository":
		var payload RepositoryPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %v", err)
		}
		return &Event{
			EventType: eventType,
			Payload:   payload,
		}, nil

	default:

		var payload any
		if err := json.Unmarshal(body, &payload); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %v", err)
		}
		return &Event{
			EventType: eventType,
			Payload:   payload,
		}, nil
	}
}
