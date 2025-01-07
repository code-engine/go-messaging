package messaging

type MessageConsumerGroup interface {
	CreateListener(name string) chan string
	ListenerExists(name string) bool
	GetListener(name string) interface{}
	Start()
	Name() string
}
