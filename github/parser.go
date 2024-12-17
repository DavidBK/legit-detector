package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	SECRET_NAME     = "GITHUB_WEBHOOK_SECRET"
	SignatureHeader = "X-Hub-Signature-256"
	EventTypeHeader = "X-GitHub-Event"
)

type EventType string

const (
	EventTypePush       EventType = "push"
	EventTypeTeam       EventType = "team"
	EventTypeRepository EventType = "repository"
)

func (e EventType) String() string {
	return string(e)
}

type Event struct {
	EventType EventType `json:"event_type"`
	Payload   any       `json:"payload"`
}

func ParseEvent(r *http.Request) (*Event, error) {
	body, err := validateRequest(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	eventType := EventType(r.Header.Get(EventTypeHeader))
	payload, err := parsePayload(eventType, body)
	if err != nil {
		return nil, err
	}

	return &Event{
		EventType: eventType,
		Payload:   payload,
	}, nil
}

func validateRequest(r *http.Request) ([]byte, error) {
	signature := r.Header.Get(SignatureHeader)
	if signature == "" {
		return nil, fmt.Errorf("missing %s header", SignatureHeader)
	}

	webhookSecret := os.Getenv(SECRET_NAME)
	if webhookSecret == "" {
		return nil, fmt.Errorf("missing %s environment variable", SECRET_NAME)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	if !isValidateSignature(body, webhookSecret, signature) {
		return nil, fmt.Errorf("invalid signature")
	}

	if r.Header.Get(EventTypeHeader) == "" {
		return nil, fmt.Errorf("missing %s header", EventTypeHeader)
	}

	return body, nil
}

func parsePayload(eventType EventType, body []byte) (any, error) {
	var payload any

	switch eventType {
	case EventTypePush:
		payload = &PushPayload{}
	case EventTypeTeam:
		payload = &TeamPayload{}
	case EventTypeRepository:
		payload = &RepositoryPayload{}
	default:
		// For unknown event types, use a generic map
		payload = &map[string]any{}
	}

	if err := json.Unmarshal(body, payload); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return payload, nil
}
