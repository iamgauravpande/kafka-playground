package producer

import (
	"context"
	"encoding/json"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
	topic  string
}

type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func NewProducer(brokers []string, topic string) *Producer {
	client, err := kgo.NewClient(kgo.SeedBrokers(brokers...))
	if err != nil {
		panic(err)
	}
	return &Producer{client: client, topic: topic}
}

func (p *Producer) SendMessage(user, message string) {
	ctx := context.Background()
	msg := Message{User: user, Message: message}
	output, _ := json.Marshal(msg)
	p.client.Produce(ctx, &kgo.Record{Topic: p.topic, Value: output}, nil)
}

func (p *Producer) Close() {
	p.client.Close()
}
