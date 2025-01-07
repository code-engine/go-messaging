package local_messaging

func NewLocalMessageChannelRepository() *LocalMessageChannelRepository {
	return &LocalMessageChannelRepository{
		Channels: make(map[string]chan string),
	}
}

type LocalMessageChannelRepository struct {
	Channels map[string]chan string
}

func (m *LocalMessageChannelRepository) CreateChannel(name string) chan string {
	if !m.ChannelExists(name) {
		m.Channels[name] = make(chan string)
		return m.Channels[name]
	} else {
		return m.Channels[name]
	}
}

func (m *LocalMessageChannelRepository) GetChannel(name string) chan string {
	return m.Channels[name]
}

func (m LocalMessageChannelRepository) ChannelExists(channelName string) bool {
	if _, ok := m.Channels[channelName]; ok {
		return true
	}

	return false
}
