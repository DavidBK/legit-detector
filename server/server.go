package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/github"
	"github.com/davidbk6/legit-detector/processor"
)

type Server struct {
	config    *configs.Config
	processor *processor.Processor
}

func NewServer(processor *processor.Processor) *Server {
	return &Server{
		config:    configs.NewConfig(),
		processor: processor,
	}
}

func (s *Server) Start() error {
	if os.Getenv(github.SECRET_NAME) == "" {
		return fmt.Errorf("missing %s environment variable", github.SECRET_NAME)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/webhook", s.handleWebhook)

	log.Printf("Starting server on port %d", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), mux)
}
