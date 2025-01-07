package local_messaging

import (
	"github.com/code-engine/go-messaging/messaging"
)

func NewLocalMessagePublisher(messageConnector messaging.MessageConnector, channelName string) *LocalMessagePublisher {
	return &LocalMessagePublisher{
		Connector:   messageConnector,
		ChannelName: channelName,
	}
}

type LocalMessagePublisher struct {
	Connector   messaging.MessageConnector
	ChannelName string
}

func (m LocalMessagePublisher) Publish(message string) {
	channel := m.Connector.GetChannel(m.ChannelName).(chan string)
	channel <- message
}
