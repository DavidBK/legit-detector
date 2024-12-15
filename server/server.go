package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/davidbk6/legit-detector/configs"
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

func handleWebhook(w http.ResponseWriter, r *http.Request, config *configs.Config) {
	log.Printf("Received webhook request")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Unsupported media type"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	fmt.Println(body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
