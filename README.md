# legit-detector

A command-line application that monitors GitHub organization activities and detects suspicious behavior using webhooks.
The application implements an extensible infrastructure for anomaly detection with current capabilities to detect:

- Code pushes during specific hours (14:00-16:00)
- Team creation with "hacker" prefix
- Repository creation and deletion within 10 minutes

## Prerequisites

- Go 1.19 or higher
- GitHub Organization admin access
- ngrok or similar tool for local webhook testing (optional)

## Configuration

Right now there is no config file so this is the the pre config values:

```bash
PORT=8080
WEBHOOK_ROUTE=/webhook
```

1. Optionally, open ngrok to expose your local server to the internet:

    ```bash
    ngrok http 8080  # Assuming your app runs on port 8080
    ```

1. Create a GitHub webhook in your organization:
   - Go to your organization settings
   - Navigate to Webhooks
   - Add webhook
   - Set Payload URL to your application server URL (e.g., `http://<ngrok-url>/webhook`)
   - Select content type: `application/json`
   - You can subscribe to all the events or select specific events
   - Set secret for signture verification

1. Create a `.env` file in the root directory with the following values:

    ```bash
    GITHUB_WEBHOOK_SECRET=<your-github-webhook-secret>
    ```

## Build and Run

```bash
go run main.go
```

## Development

The project is written in a pub sub architecture.
Each `detector` subscribes to the events he is interested in and the `dispatcher` sends the events to them when they arrive.
When a `detector` detects an anomaly it sends a notification using `notifier`.
The detectors rules can be complex and built from multiple events from github.

### Project components

The `server` is listening to the incoming webhooks and pass them to the `dispatcher`.

The `github` is responsible for the github types and the webhook verification.

The `dispatcher` is responsible for sending the events to the rules.

The `detectors` are responsible for detecting anomalies and sending alert using the `notifications`.

### Add new notification methods

Notifications are implemented in the `notifications` package.

To add a new notification method you need to implement the `Notifier` interface:

```go
type Notifier interface {
	Notify(event Event) error
}
```

If you want all The existing detection rules to use the new notifier you need to add it to the curr `NotificationManager` in the `main.go` file,
for example:

```go
func createNotifiers(n *notifications.NotificationManager) {
	n.AddNotifier(notifications.NewLogNotifier())
	n.AddNotifier(notifications.<YourNewNotifier>)
}
```

### Add new detection rules

Detecting rules are implemented in the `detectors` package.

To add a new detection rule you need to implement the `Detector` interface:

```go
type Detector interface {
	GetEventTypes() []github.EventType
	Handle(*github.Event)
}
```

The `GetEventTypes` method should return the event types the detector is interested in.

The `Handle` method will be called when the event arrives,
so if the detection rule is build from multiple events you need to handle internally the state.
You can call `notifier.Notify` when you detect an anomaly.

To register the new detector you need to add it to the `dispatcher` in the `main.go` file,
for example:

```go
func registerRules(ed *events.EventDispatcher, n *notifications.NotificationManager) {
	ed.Subscribe(detectors.NewPushTimeRule(n))
	ed.Subscribe(detectors.NewTeamNameRule(n))
	ed.Subscribe(detectors.NewRepoLifeTimeRule(n))
	ed.Subscribe(detectors.<YourNewRule>(n))
}
```

## TODO

I didn't have time to implement all the features I wanted.

There is [TODO.md](./TODO.md) file with the next steps.
