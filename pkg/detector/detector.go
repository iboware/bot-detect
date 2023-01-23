package detector

import "context"

type Detector interface {
	Detect(ctx context.Context, message []byte) (bool, error)
}
