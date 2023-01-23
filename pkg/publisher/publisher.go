package publisher

import "context"

type Publisher interface {
	Publish(ctx context.Context, topicID string, msg []byte) error
}
