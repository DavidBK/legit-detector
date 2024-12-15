package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/github"
)

func CreateServer() {
	config := configs.NewConfig()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		handleHealth(w, r, config)
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		handleWebhook(w, r, config)
	})

	log.Printf("Starting server on port %d", config.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request, _ *configs.Config) {
	log.Printf("Received health check request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	log.Printf("Sent health check response")
}

func handleWebhook(w http.ResponseWriter, r *http.Request, _ *configs.Config) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	event, err := github.ParseEvent(r)
	if err != nil {
		log.Printf("Failed to parse event: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	eventType := event.EventType
	log.Printf("Received %s event", eventType)

	switch eventType {
	case "push":
		handlePushEvent(event.Payload)
	case "pull_request":
		handlePullRequestEvent(event.Payload)
	default:
		log.Printf("Unhandled event type: %s", eventType)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"event":  eventType,
	})
}

func handlePushEvent(payload interface{}) {
	log.Printf("Received push event: %v", payload)
}

func handlePullRequestEvent(payload interface{}) {
	log.Printf("Received pull event: %v", payload)
}
