package notifications

import "fmt"

type LogNotifier struct{}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{}
}

const notificationFormat = "!!! ALERT !!!\n" +
	"\tTitle: %s\n" +
	"\tOrganization: %s\n" +
	"\tTimestamp: %s\n" +
	"\tMessage: %s\n" +
	"!!!!!!!!!!!!!\n"

func (n *LogNotifier) Notify(notification Notification) {
	fmt.Printf(notificationFormat, notification.Title, notification.Organization, notification.Timestamp, notification.Message)
}
