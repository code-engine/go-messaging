package local_messaging

import (
	"log/slog"

	"github.com/code-engine/go-messaging/messaging"
)

func NewLocalMessageSubscriber(messageConnector messaging.MessageConnector, handler messaging.MessageHandler, channelName string) *LocalMessageSubscriber {
	return &LocalMessageSubscriber{
		MessageConnector: messageConnector,
		Handler:          handler,
		ChannelName:      channelName,
	}
}

type LocalMessageSubscriber struct {
	MessageConnector messaging.MessageConnector
	Handler          messaging.MessageHandler
	ChannelName      string
}

func (m LocalMessageSubscriber) Subscribe() {
	slog.Info("Subscribing", "channelName", m.ChannelName)
	channel := m.MessageConnector.GetChannel(m.ChannelName).(chan string)

	for {
		message := <-channel
		m.Handler(message)
	}
}

func (m LocalMessageSubscriber) Configure() {}
