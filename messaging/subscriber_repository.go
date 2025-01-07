package messaging

func NewMessageSubscriberRepository() MessageSubscriberRepository {
	return MessageSubscriberRepository{
		MessageSubscribers: make(map[string]MessageSubscriber),
	}
}

type MessageSubscriberRepository struct {
	MessageSubscribers map[string]MessageSubscriber
}

func (m MessageSubscriberRepository) GetMessageSubscribers() map[string]MessageSubscriber {
	return m.MessageSubscribers
}

func (m *MessageSubscriberRepository) AddMessageSubscriber(messageSubscriberName string, messageSubscriber MessageSubscriber) {
	m.MessageSubscribers[messageSubscriberName] = messageSubscriber
}
