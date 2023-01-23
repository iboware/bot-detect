package detector

import "context"

//go:generate mockgen -package mock -destination=./mock/detector_mock.go . Detector
type Detector interface {
	Detect(ctx context.Context, message []byte) (bool, error)
}
