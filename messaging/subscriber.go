package messaging

type MessageSubscriber interface {
	Configure()
	Subscribe()
}
