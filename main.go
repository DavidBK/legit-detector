package main

import (
	"log"

	"github.com/davidbk6/legit-detector/detectors"
	"github.com/davidbk6/legit-detector/notifications"
	"github.com/davidbk6/legit-detector/processor"
	"github.com/davidbk6/legit-detector/registry"
	"github.com/davidbk6/legit-detector/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	createNotifiers()
	reg := registry.NewRegistry()
	registerRules(reg)

	proc := processor.NewProcessor(reg)
	srv := server.NewServer(proc)
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
		panic(err)
	}
}

func createNotifiers() {
	notificationManager := notifications.GetManager()
	notificationManager.AddNotifier(notifications.NewLogNotifier())
}

func registerRules(registry *registry.Registry) {
	registry.Register(detectors.NewPushTimeRule())
	registry.Register(detectors.NewTeamNameRule())
	registry.Register(detectors.NewRepoLifeTimeRule())
}
