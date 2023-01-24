package protoutil

import (
	model "github.com/iboware/bot-detect/gen/proto/go"
	"google.golang.org/protobuf/proto"
)

func IncomingMsgFromProto(message []byte) (*model.IncomingMsg, error) {
	messagePb := &model.IncomingMsg{}
	if err := proto.Unmarshal(message, messagePb); err != nil {
		return nil, err
	}
	return messagePb, nil
}

func OutgoingMsgToByteArray(message *model.OutgoingMsg) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	if data, err = proto.Marshal(message); err != nil {
		return nil, err
	}
	return data, nil
}

func IncomingMsgToByteArray(message *model.IncomingMsg) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	if data, err = proto.Marshal(message); err != nil {
		return nil, err
	}
	return data, nil
}
