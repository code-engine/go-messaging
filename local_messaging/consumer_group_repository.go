package local_messaging

import (
	"log/slog"

	"github.com/code-engine/go-messaging/messaging"
)

func NewLocalMessageConsumerGroupRepository() *LocalMessageConsumerGroupRepository {
	return &LocalMessageConsumerGroupRepository{
		ConsumerGroups: make(map[string]messaging.MessageConsumerGroup),
	}
}

type LocalMessageConsumerGroupRepository struct {
	ConsumerGroups map[string]messaging.MessageConsumerGroup
}

func (l *LocalMessageConsumerGroupRepository) ConsumerGroupExists(consumerGroupName string) bool {
	if _, ok := l.ConsumerGroups[consumerGroupName]; ok {
		return true
	}

	return false
}

func (l *LocalMessageConsumerGroupRepository) CreateConsumerGroup(consumerGroupName string, channel chan string) {
	slog.Info("Creating Consumer Group", "consumer_group_name", consumerGroupName)
	l.ConsumerGroups[consumerGroupName] = NewLocalMessageConsumerGroup(consumerGroupName, channel)
}

func (l *LocalMessageConsumerGroupRepository) GetConsumerGroup(name string) messaging.MessageConsumerGroup {
	consumerGroup := l.ConsumerGroups[name]
	return consumerGroup
}

func (l *LocalMessageConsumerGroupRepository) AddConsumerGroup(consumerGroupName string, consumerGroup messaging.MessageConsumerGroup) {
	l.ConsumerGroups[consumerGroupName] = consumerGroup
}
