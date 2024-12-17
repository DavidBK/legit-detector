package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	SECRET_NAME     = "GITHUB_WEBHOOK_SECRET"
	SignatureHeader = "X-Hub-Signature-256"
	EventTypeHeader = "X-GitHub-Event"
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
		Name      string    `json:"name"`
		Id        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"repository"`
}

func ParseEvent(r *http.Request) (*Event, error) {
	body, err := validateRequest(r)

	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}

	eventType := r.Header.Get(EventTypeHeader)
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

func parsePayload(eventType string, body []byte) (any, error) {
	var payload any

	switch eventType {
	case "push":
		payload = &PushPayload{}
	case "team":
		payload = &TeamPayload{}
	case "repository":
		payload = &RepositoryPayload{}
	default:
		payload = &map[string]any{}
	}

	if err := json.Unmarshal(body, payload); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return payload, nil
}
