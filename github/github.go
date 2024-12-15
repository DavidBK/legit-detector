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

func ParseEvent(r *http.Request) (*Event, error) {
	eventType := r.Header.Get("X-GitHub-Event")
	if eventType == "" {
		return nil, fmt.Errorf("missing X-GitHub-Event header")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &Event{
		EventType: eventType,
		Payload:   payload,
	}, nil
}
