package messaging

func NewMessagePublisherRepository() MessagePublisherRepository {
	return MessagePublisherRepository{
		MessagePublishers: make(map[string]MessagePublisher),
	}
}

type MessagePublisherRepository struct {
	MessagePublishers map[string]MessagePublisher
}

func (m MessagePublisherRepository) GetMessagePublishers() map[string]MessagePublisher {
	return m.MessagePublishers
}

func (m *MessagePublisherRepository) AddMessagePublisher(messagePublisherName string, messagePublisher MessagePublisher) {
	m.MessagePublishers[messagePublisherName] = messagePublisher
}
