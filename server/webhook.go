package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidbk6/legit-detector/github"
)

func (s *Server) handleWebhook(w http.ResponseWriter, r *http.Request) {
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
		err := s.handleEvent(event)
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

func (s *Server) handleEvent(event *github.Event) error {
	log.Printf("Processing %s event", event.EventType)
	err := s.eventDispatcher.Dispatch(event)
	if err != nil {
		return err
	}
	log.Printf("Finished processing %s event", event.EventType)
	return nil
}
