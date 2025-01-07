package local_messaging

import "github.com/code-engine/go-messaging/messaging"

func NewLocalMessageBroadcastSubscriber(messageConnector messaging.MessageConnector, handler messaging.MessageHandler, listenerName, consumerGroupName string, channel chan string) *LocalMessageBroadcastSubscriber {
	return &LocalMessageBroadcastSubscriber{
		MessageConnector:  messageConnector,
		Handler:           handler,
		ListenerName:      listenerName,
		ConsumerGroupName: consumerGroupName,
		Channel:           channel,
	}
}

type LocalMessageBroadcastSubscriber struct {
	MessageConnector  messaging.MessageConnector
	Handler           messaging.MessageHandler
	ListenerName      string
	ConsumerGroupName string
	Channel           chan string
	Listener          chan string
}

func (m *LocalMessageBroadcastSubscriber) Configure() {
	m.Listener = m.MessageConnector.CreateListener(m.ListenerName, m.ConsumerGroupName, m.Channel).(chan string)
}

func (m LocalMessageBroadcastSubscriber) Subscribe() {
	for {
		message := <-m.Listener
		m.Handler(message)
	}
}
