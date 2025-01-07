package local_messaging

import (
	"encoding/json"
	"time"

	utils "github.com/code-engine/go-utils"

	"github.com/code-engine/go-messaging/messaging"
)

func NewLocalMessageGenerator() *LocalMessageGenerator {
	return &LocalMessageGenerator{Wait: 2 * time.Second}
}

type LocalMessage struct {
	ID        string    `json:"id"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

type LocalMessageGenerator struct {
	Wait time.Duration
}

func (m LocalMessageGenerator) Generate(publisher messaging.MessagePublisher) {
	messages := []LocalMessage{
		{ID: utils.NewUUID(), Event: "UserCreated", Timestamp: time.Now()},
		{ID: utils.NewUUID(), Event: "UserDestroyed", Timestamp: time.Now()},
		{ID: utils.NewUUID(), Event: "ProjectCreated", Timestamp: time.Now()},
		{ID: utils.NewUUID(), Event: "ProjectDestroyed", Timestamp: time.Now()},
	}

	for _, message := range messages {
		jsonMessage, _ := json.Marshal(message)
		publisher.Publish(string(jsonMessage))
		time.Sleep(m.Wait)
	}
}
