package consumer

import "context"

//go:generate mockgen -package mock -destination=./mock/consumer_mock.go . Consumer
type Consumer interface {
	Start(ctx context.Context) error
}
