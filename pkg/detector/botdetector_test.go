package detector

import (
	"context"
	"testing"

	"github.com/iboware/bot-detect/pkg/publisher"
)

func TestBotDetector_Detect(t *testing.T) {
	type fields struct {
		Publisher       publisher.Publisher
		IPBlockList     map[string]string
		OutgoingTopicId string
	}
	type args struct {
		ctx     context.Context
		message []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &BotDetector{
				Publisher:   tt.fields.Publisher,
				IPBlockList: tt.fields.IPBlockList,
			}
			got, err := d.Detect(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("BotDetector.Detect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BotDetector.Detect() = %v, want %v", got, tt.want)
			}
		})
	}
}
