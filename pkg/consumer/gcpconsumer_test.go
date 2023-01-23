package consumer

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/iboware/bot-detect/pkg/detector"
)

func TestGCPConsumer_Start(t *testing.T) {
	type fields struct {
		Client       *pubsub.Client
		Subscription *pubsub.Subscription
		Detector     detector.Detector
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &GCPConsumer{
				Client:       tt.fields.Client,
				Subscription: tt.fields.Subscription,
				Detector:     tt.fields.Detector,
			}
			if err := ps.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GCPConsumer.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
