package notifications

import (
	"sync"
	"time"
)

type Notification struct {
	Message      string
	EventType    string
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

var (
	manager *NotificationManager
	once    sync.Once
)

func GetManager() *NotificationManager {
	once.Do(func() {
		manager = &NotificationManager{
			notifiers: make([]Notifier, 0),
		}
	})
	return manager
}

func (m *NotificationManager) AddNotifier(notifier Notifier) {
	m.notifiers = append(m.notifiers, notifier)
}

func (m *NotificationManager) NotifyAll(notification Notification) {
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
