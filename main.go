package main

import (
	"github.com/davidbk6/legit-detector/notifications"
	"github.com/davidbk6/legit-detector/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	notificationManager := notifications.GetManager()
	notificationManager.AddNotifier(notifications.NewLogNotifier())

	server.CreateServer()
}
