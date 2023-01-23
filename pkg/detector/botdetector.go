package detector

import (
	"context"
	"strings"

	model "github.com/iboware/bot-detect/gen/proto/go"
	"github.com/iboware/bot-detect/pkg/protoutil"
	"github.com/iboware/bot-detect/pkg/publisher"
	"github.com/iboware/bot-detect/pkg/timeutil"
)

const CONVERSION_WORD = "thank-you"

type BotDetector struct {
	Publisher       publisher.Publisher
	IPBlockList     map[string]string
	OutgoingTopicId string
}

func NewDetector(publisher publisher.Publisher, ipBlockList map[string]string, outgoingTopicId string) Detector {
	return &BotDetector{
		Publisher:       publisher,
		IPBlockList:     ipBlockList,
		OutgoingTopicId: outgoingTopicId,
	}
}

func (d *BotDetector) Detect(ctx context.Context, message []byte) (bool, error) {

	in, err := protoutil.IncomingMsgFromProto(message)
	if err != nil {
		return false, err
	}

	ts := timeutil.TryParseUnixTime(in.Timestamp)

	out := model.OutgoingMsg{
		IsBot:        false,
		IsConversion: false,
		Key:          in.Key,
		Timestamp:    timeutil.FormatTimeP(ts),
	}

	// Check if ip block list contains the ip address of the request.
	if _, out.IsBot = d.IPBlockList[in.Ip]; !out.IsBot {
		// ip address is not in block list, check if the request is a conversion.
		out.IsConversion = strings.Contains(in.Url, CONVERSION_WORD)
	}

	outData, err := protoutil.OutgoingMsgToByteArray(&out)
	if err != nil {
		return false, err
	}

	if err := d.Publisher.Publish(ctx, d.OutgoingTopicId, outData); err != nil {
		return false, err
	}

	return true, nil
}
