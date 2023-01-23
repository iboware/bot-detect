package publisher

import "context"

//go:generate mockgen -package mock -destination=./mock/publisher_mock.go . Publisher
type Publisher interface {
	Publish(ctx context.Context, topicID string, msg []byte) error
}
