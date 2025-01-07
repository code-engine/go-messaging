package local_messaging

import "github.com/code-engine/go-messaging/messaging"

func NewLocalMessageAdapter() *LocalMessageAdapter {
	return &LocalMessageAdapter{
		MessageConnector: NewLocalMessageConnector(),
	}
}

type LocalMessageAdapter struct {
	MessageConnector messaging.MessageConnector
}

func (as *LocalMessageAdapter) CreateMessageSubscriber(messagingHandler messaging.MessageHandler, channelName string) messaging.MessageSubscriber {
	// Test this
	if !as.MessageConnector.ChannelExists(channelName) {
		as.MessageConnector.CreateChannel(channelName)
	}

	return NewLocalMessageSubscriber(as.MessageConnector, messagingHandler, channelName)
}

func (as *LocalMessageAdapter) CreateMessageBroadcastSubscriber(messageHandler messaging.MessageHandler, listenerName, consumerGroupName, channelName string) messaging.MessageSubscriber {
	var channel chan string
	if as.MessageConnector.ChannelExists(channelName) {
		channel = as.MessageConnector.GetChannel(channelName).(chan string)
	} else {
		channel = as.MessageConnector.CreateChannel(channelName).(chan string)
	}

	return NewLocalMessageBroadcastSubscriber(as.MessageConnector, messageHandler, listenerName, consumerGroupName, channel)
}

func (as *LocalMessageAdapter) CreateMessagePublisher(channelName string) messaging.MessagePublisher {
	// Test this
	if !as.MessageConnector.ChannelExists(channelName) {
		as.MessageConnector.CreateChannel(channelName)
	}

	return NewLocalMessagePublisher(as.MessageConnector, channelName)
}

func (as *LocalMessageAdapter) Start() {
	as.MessageConnector.Start()
}
