package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/events"
	"github.com/davidbk6/legit-detector/github"
)

type Server struct {
	config          *configs.Config
	eventDispatcher *events.EventDispatcher
}

func NewServer(c *configs.Config, ed *events.EventDispatcher) *Server {
	return &Server{
		config:          c,
		eventDispatcher: ed,
	}
}

func (s *Server) Start() error {
	if os.Getenv(github.SECRET_NAME) == "" {
		return fmt.Errorf("missing %s environment variable", github.SECRET_NAME)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/webhook", s.handleWebhook)

	log.Printf("Starting server on port %d", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), mux)
}
