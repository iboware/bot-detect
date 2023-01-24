package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type GCPPublisher struct {
	Client *pubsub.Client
	Topic  *pubsub.Topic
}

func NewGCPPublisher(client *pubsub.Client, topicID string) Publisher {
	return &GCPPublisher{
		Client: client,
		Topic:  client.Topic(topicID),
	}
}

func (p *GCPPublisher) Publish(ctx context.Context, msg []byte) error {
	result := p.Topic.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}
