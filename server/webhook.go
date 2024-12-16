package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/processor"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	event, err := github.ParseEvent(r)
	if err != nil {
		log.Printf("Failed to parse event: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	eventType := event.EventType
	log.Printf("Received %s event", eventType)

	go func() {
		err := processor.HandleEvent(event)
		if err != nil {
			log.Printf("Failed to process event: %v", err)
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"event":  eventType,
	})
}
