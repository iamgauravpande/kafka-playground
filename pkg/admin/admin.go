package admin

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Admin struct {
	client *kadm.Client
}

// Create New Admin Client from kafka client:

func NewAdmin(brokers []string) *Admin {
	client, err := kgo.NewClient(kgo.SeedBrokers(brokers...)) // returns new kafka client
	if err != nil {
		panic(err)
	}
	admin := kadm.NewClient(client) // returns an admin Client
	return &Admin{client: admin}
}

// List the topics created:

func (a *Admin) TopicExist(topic string) bool {
	ctx := context.Background()
	topicDetails, err := a.client.ListTopics(ctx) // return topic details
	if err != nil {
		panic(err)
	}
	for _, metadata := range topicDetails { // iterate over to get topic name and print
		if metadata.Topic == topic {
			fmt.Println("Topic already exist")
			return true
		}
	}
	fmt.Println("Topic doesnt exist!")
	return false
}

func (a *Admin) TopicCreate(topic string) {
	ctx := context.Background()
	resp, err := a.client.CreateTopics(ctx, 3, 1, nil, topic)
	if err != nil {
		panic(err)
	}
	for _, ctr := range resp {
		if ctr.Err != nil {
			fmt.Println("Unable to create Topic", ctr.Topic, ctr.Err)
		} else {
			fmt.Println("Topic created!", ctr.Topic)
		}
	}
}

func (a *Admin) Close() {
	a.client.Close()
}
