package consumer

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/iboware/bot-detect/pkg/detector"
)

type GCPConsumer struct {
	Client       *pubsub.Client
	Subscription *pubsub.Subscription
	Detector     detector.Detector
}

func NewGCPConsumer(client *pubsub.Client, subscription *pubsub.Subscription, detector detector.Detector) Consumer {
	return &GCPConsumer{
		Client:       client,
		Subscription: subscription,
		Detector:     detector,
	}
}

func (ps *GCPConsumer) Start(ctx context.Context) error {
	// Receive blocks until the context is cancelled or an error occurs.
	err := ps.Subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if ok, err := ps.Detector.Detect(ctx, msg.Data); !ok {
			fmt.Printf("error occurred while detecting: %s", err.Error())
			msg.Nack()
			return
		}
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("consumer returned error: %v", err)
	}
	return nil
}
