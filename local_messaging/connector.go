package local_messaging

import (
	"log/slog"

	"github.com/code-engine/go-messaging/messaging"
)

func NewLocalMessageConnector() *LocalMessageConnector {
	return &LocalMessageConnector{
		ChannelRepository:       NewLocalMessageChannelRepository(),
		ConsumerGroupRepository: NewLocalMessageConsumerGroupRepository(),
	}
}

type LocalMessageConnector struct {
	ChannelRepository       *LocalMessageChannelRepository
	ConsumerGroupRepository *LocalMessageConsumerGroupRepository
}

func (m LocalMessageConnector) ChannelExists(channelName string) bool {
	return m.ChannelRepository.ChannelExists(channelName)
}

func (m LocalMessageConnector) ConsumerGroupExists(consumerGroupName string) bool {
	return m.ConsumerGroupRepository.ConsumerGroupExists(consumerGroupName)
}

func (m LocalMessageConnector) ListenerExists(listenerName, consumerGroupName string) bool {
	consumerGroup := m.ConsumerGroupRepository.GetConsumerGroup(consumerGroupName)
	return consumerGroup.ListenerExists(listenerName)
}

func (m *LocalMessageConnector) CreateChannel(channelName string) interface{} {
	// Check if channel exists
	return m.ChannelRepository.CreateChannel(channelName)
}

func (m *LocalMessageConnector) CreateConsumerGroup(consumerGroupName, channelName string) messaging.MessageConsumerGroup {
	// Check if channel exists
	channel := m.ChannelRepository.GetChannel(channelName)
	m.ConsumerGroupRepository.CreateConsumerGroup(consumerGroupName, channel)
	return m.ConsumerGroupRepository.GetConsumerGroup(consumerGroupName)
}

func (m *LocalMessageConnector) CreateListener(listenerName, consumerGroupName string, channel chan string) interface{} {
	var consumerGroup messaging.MessageConsumerGroup
	if !m.ConsumerGroupRepository.ConsumerGroupExists(consumerGroupName) {
		consumerGroup = NewLocalMessageConsumerGroup(consumerGroupName, channel)
		m.ConsumerGroupRepository.AddConsumerGroup(consumerGroupName, consumerGroup)
	} else {
		consumerGroup = m.ConsumerGroupRepository.GetConsumerGroup(consumerGroupName)
	}

	listener := consumerGroup.CreateListener(listenerName)

	return listener
}

func (m LocalMessageConnector) GetChannel(channelName string) interface{} {
	// Check if channel exists
	return m.ChannelRepository.GetChannel(channelName)
}

func (m *LocalMessageConnector) GetConsumerGroup(consumerGroupName string) messaging.MessageConsumerGroup {
	return m.ConsumerGroupRepository.GetConsumerGroup(consumerGroupName)
}

func (m LocalMessageConnector) GetListener(listenerName, consumerGroupName string) interface{} {
	consumerGroup := m.ConsumerGroupRepository.GetConsumerGroup(consumerGroupName)
	listener := consumerGroup.GetListener(listenerName)
	return listener
}

func (m *LocalMessageConnector) Start() {
	slog.Info("Starting Message Connector")
	m.StartConsumerGroups()
}

func (m *LocalMessageConnector) StartConsumerGroups() {
	slog.Info("Starting ConsumerGroups", "length", len(m.ConsumerGroupRepository.ConsumerGroups))

	for _, consumerGroup := range m.ConsumerGroupRepository.ConsumerGroups {
		slog.Info("Starting ConsumerGroup", "consumer_group", consumerGroup.Name())
		// Manage goroutines better
		go consumerGroup.Start()
	}
}
