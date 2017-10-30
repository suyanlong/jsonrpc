package libproto

import (
	"jsonrpc/snappy"
	"github.com/golang/protobuf/proto"
)

func CmdId(submodule, topic uint32) uint32 {
	return (submodule << 16) + topic
}

func DeCmdId(cmdId uint32) (uint32, uint32) {
	submodule := cmdId >> 16
	var sub = submodule
	var tmp = submodule << 16
	topic := cmdId - tmp
	return sub, topic
}

const ZEROORIGIN uint32 = 99999

func CreateMsg(sub, topic uint32, msgType MsgType, content []byte) ([]byte, error) {
	msg := Message{
		CmdId:   CmdId(sub, topic),
		Type:    msgType,
		Operate: OperateType_BROADCAST,
		Origin:  ZEROORIGIN,
		//compress data
		Content: snappy.CitaCompress(content),
	}
	return proto.Marshal(&msg)
}

func CreateMsgEx(sub, topic, origin uint32, msgType MsgType, operate OperateType, content []byte) ([]byte, error) {
	msg := Message{
		CmdId:   CmdId(sub, topic),
		Type:    msgType,
		Operate: operate,
		Origin:  origin,
		//compress data
		Content: snappy.CitaCompress(content),
	}

	return proto.Marshal(&msg)
}

func ParseMsg(data []byte) (uint32, MsgType, proto.Message, error) {
	var pb Message
	if err := proto.Unmarshal(data, &pb); err != nil {
		return 0, 0, nil, err
	} else {
		deData := snappy.CitaDecompress(pb.Content)
		var content proto.Message
		if pb.Type == MsgType_RESPONSE {
			content = new(Response)
		}
		if err := proto.Unmarshal(deData, content); err != nil {
			return 0, 0, nil, err
		} else {
			return pb.CmdId, pb.Type, content, nil
		}

	}
}
