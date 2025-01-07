package messaging

type MessageConnector interface {
	ChannelExists(channelName string) bool
	ConsumerGroupExists(consumerGroupName string) bool
	ListenerExists(consumerGroupName, listenerName string) bool

	CreateChannel(channelName string) interface{}
	CreateConsumerGroup(consumerGroupName, channelName string) MessageConsumerGroup
	CreateListener(listenerName, consumerGroupName string, channel chan string) interface{}

	GetChannel(channelName string) interface{}
	GetConsumerGroup(consumerGroupName string) MessageConsumerGroup
	GetListener(listenerName, consumerGroupName string) interface{}

	Start()
}
