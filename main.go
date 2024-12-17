package main

import (
	"log"

	"github.com/davidbk6/legit-detector/configs"
	"github.com/davidbk6/legit-detector/detectors"
	"github.com/davidbk6/legit-detector/events"
	"github.com/davidbk6/legit-detector/notifications"
	"github.com/davidbk6/legit-detector/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	notifer := notifications.NewNotificationManager()
	createNotifiers(notifer)

	eventDispatcher := events.NewEventDispatcher()
	registerRules(eventDispatcher, notifer)

	config := configs.NewConfig()

	srv := server.NewServer(config, eventDispatcher)
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
		panic(err)
	}
}

func createNotifiers(n *notifications.NotificationManager) {
	n.AddNotifier(notifications.NewLogNotifier())
}

func registerRules(ed *events.EventDispatcher, n *notifications.NotificationManager) {
	ed.Subscribe(detectors.NewPushTimeRule(n))
	ed.Subscribe(detectors.NewTeamNameRule(n))
	ed.Subscribe(detectors.NewRepoLifeTimeRule(n))
}
