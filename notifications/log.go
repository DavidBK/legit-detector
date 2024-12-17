package notifications

import "fmt"

type LogNotifier struct{}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{}
}

const notificationFormat = "!!! ALERT !!!\n" +
	"\tEvent: %s\n" +
	"\tOrganization: %s\n" +
	"\tTimestamp: %s\n" +
	"\tMessage: %s\n" +
	"!!!!!!!!!!!!!!!\n"

func (n *LogNotifier) Notify(notification Notification) {
	fmt.Printf(notificationFormat, notification.EventType, notification.Organization, notification.Timestamp, notification.Message)
}
