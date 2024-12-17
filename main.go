package main

import (
	"log"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/detectors"
	"github.com/davidbk6/legit-detector/event"
	"github.com/davidbk6/legit-detector/notifications"
	"github.com/davidbk6/legit-detector/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	createNotifiers()
	eventDispatcher := event.NewEventDispatcher()
	registerRules(eventDispatcher)

	config := configs.NewConfig()

	srv := server.NewServer(config, eventDispatcher)
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
		panic(err)
	}
}

func createNotifiers() {
	notificationManager := notifications.GetManager()
	notificationManager.AddNotifier(notifications.NewLogNotifier())
}

func registerRules(ed *event.EventDispatcher) {
	ed.Subscribe(detectors.NewPushTimeRule())
	ed.Subscribe(detectors.NewTeamNameRule())
	ed.Subscribe(detectors.NewRepoLifeTimeRule())
}
