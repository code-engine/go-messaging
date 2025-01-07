package messaging

type MessageSubscription struct {
	Name              string
	Handler           func(message string)
	ChannelName       string
	ConsumerGroupName string
	ListenerName      string
}
