package publisher

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
)

func TestGCPPublisher_Publish(t *testing.T) {
	type fields struct {
		Client *pubsub.Client
		Topic  *pubsub.Topic
	}
	type args struct {
		ctx context.Context
		msg []byte
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
			p := &GCPPublisher{
				Client: tt.fields.Client,
				Topic:  tt.fields.Topic,
			}
			if err := p.Publish(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("GCPPublisher.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
