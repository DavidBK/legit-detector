package notifications

import (
	"sync"
	"time"
)

type Notification struct {
	Message      string
	Title        string
	Organization string
	Timestamp    time.Time
	Metadata     map[string]any
}

type Notifier interface {
	Notify(notification Notification)
}

type NotificationManager struct {
	notifiers []Notifier
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{}
}

func (m *NotificationManager) AddNotifier(notifier Notifier) {
	m.notifiers = append(m.notifiers, notifier)
}

func (m *NotificationManager) Notify(notification Notification) {
	wg := sync.WaitGroup{}
	for _, notifier := range m.notifiers {
		wg.Add(1)
		go func(n Notifier) {
			defer wg.Done()
			n.Notify(notification)
		}(notifier)
	}

	wg.Wait()
}
