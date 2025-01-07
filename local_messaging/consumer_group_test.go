package local_messaging_test

import (
	"testing"

	"github.com/code-engine/go-messaging/local_messaging"
	"github.com/stretchr/testify/assert"
)

func TestCreateListener(t *testing.T) {
	name := "local-consumer-group"
	channelReference := make(chan string)
	localMessageConsumerGroup := local_messaging.NewLocalMessageConsumerGroup(name, channelReference)

	listenerName := "test-listener"
	localMessageConsumerGroup.CreateListener(listenerName)

	assert.NotNil(t, localMessageConsumerGroup.Listeners[listenerName])
}

func TestListenerExistsReturnsTrueWhenItDoesExist(t *testing.T) {
	name := "local-consumer-group"
	channelReference := make(chan string)
	localMessageConsumerGroup := local_messaging.NewLocalMessageConsumerGroup(name, channelReference)

	listenerName := "test-listener"
	localMessageConsumerGroup.CreateListener(listenerName)

	assert.True(t, localMessageConsumerGroup.ListenerExists(listenerName), "expected true")
}

func TestListenerExistsReturnsFalseWhenItDoesNotExist(t *testing.T) {
	name := "local-consumer-group"
	channelReference := make(chan string)
	localMessageConsumerGroup := local_messaging.NewLocalMessageConsumerGroup(name, channelReference)

	listenerName := "test-listener"

	assert.False(t, localMessageConsumerGroup.ListenerExists(listenerName), "expected false")
}

func TestConsumerGroupStart(t *testing.T) {
	name := "local-consumer-group"
	channelReference := make(chan string)
	localMessageConsumerGroup := local_messaging.NewLocalMessageConsumerGroup(name, channelReference)

	listener1 := localMessageConsumerGroup.CreateListener("listener1")
	listener2 := localMessageConsumerGroup.CreateListener("listener2")
	listener3 := localMessageConsumerGroup.CreateListener("listener3")

	// Manage goroutines better
	go localMessageConsumerGroup.Start()

	testMessage := "test"
	channelReference <- testMessage

	listenerMsg1 := <-listener1
	listenerMsg2 := <-listener2
	listenerMsg3 := <-listener3

	assert.Equal(t, listenerMsg1, testMessage, "expected messages to be equal")
	assert.Equal(t, listenerMsg2, testMessage, "expected messages to be equal")
	assert.Equal(t, listenerMsg3, testMessage, "expected messages to be equal")
}
