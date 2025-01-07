package messaging

type MessageAdapter interface {
	CreateMessageSubscriber(messagingHandler MessageHandler, channelName string) MessageSubscriber
	CreateMessageBroadcastSubscriber(messageHandler MessageHandler, listenerName, consumerGroupName, channelName string) MessageSubscriber
	CreateMessagePublisher(channelName string) MessagePublisher
	Start()
}
