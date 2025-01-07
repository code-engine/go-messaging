package local_messaging

import "log/slog"

func NewLocalMessageConsumerGroup(name string, channelReference chan string) *LocalMessageConsumerGroup {
	slog.Info("Creating ConsumerGroup", "name", name)
	return &LocalMessageConsumerGroup{
		ChannelReference: channelReference,
		Listeners:        make(map[string]chan string),
		name:             name,
	}
}

type LocalMessageConsumerGroup struct {
	ChannelReference chan string
	Listeners        map[string]chan string
	name             string
}

func (m *LocalMessageConsumerGroup) CreateListener(name string) chan string {
	slog.Info("Creating listener", "listener", name)
	if !m.ListenerExists(name) {
		m.Listeners[name] = make(chan string)
	}

	return m.Listeners[name]
}

func (m LocalMessageConsumerGroup) ListenerExists(name string) bool {
	if _, ok := m.Listeners[name]; ok {
		return true
	}

	return false
}

func (m *LocalMessageConsumerGroup) GetListener(listenerName string) interface{} {
	return m.Listeners[listenerName]
}

func (m *LocalMessageConsumerGroup) Start() {
	slog.Info("Starting consumer group listeners")
	// Make this better
	for {
		message := <-m.ChannelReference

		for _, listener := range m.Listeners {
			listener <- message
		}
	}
}

func (m LocalMessageConsumerGroup) Name() string {
	return m.name
}
