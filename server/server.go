package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/processor"
)

func CreateServer() {
	config := configs.NewConfig()

	// TODO: handle dinamic secret
	if os.Getenv(github.SECRET_NAME) == "" {
		panic(fmt.Sprintf("missing %s environment variable", github.SECRET_NAME))
	}

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
