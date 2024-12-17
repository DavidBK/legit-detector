package server

import (
	"log"
	"net/http"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received health check request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	log.Printf("Sent health check response")
}
