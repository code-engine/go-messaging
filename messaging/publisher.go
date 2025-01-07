package messaging

type MessagePublisher interface {
	Publish(message string)
}
