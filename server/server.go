package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/github"
)

func CreateServer() {
	config := configs.NewConfig()

	// TODO: handle dynamic secret
	if os.Getenv(github.SECRET_NAME) == "" {
		panic(fmt.Sprintf("missing %s environment variable", github.SECRET_NAME))
	}

	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/webhook", handleWebhook)

	log.Printf("Starting server on port %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
