package main

import (
	"errors"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type MsgHandler func(msgid uint16, msg interface{})

type MessageInfo struct {
	msgType    reflect.Type
	msgHandler MsgHandler
}

var (
	msg_map = make(map[uint16]*MessageInfo)
)

func RegisterMessage(msgid uint16, msg interface{}, handler MsgHandler) {
	info := new(MessageInfo)
	info.msgType = reflect.TypeOf(msg.(proto.Message))
	info.msgHandler = handler

	msg_map[msgid] = info
}

func HandleRawData(msgid uint16, data []byte) error {
	if info, ok := msg_map[msgid]; ok {
		msg := reflect.New(info.msgType.Elem()).Interface()
		err := proto.Unmarshal(data, msg.(proto.Message))
		if err != nil {
			return err
		}
		info.msgHandler(msgid, msg)
		return err
	}
	return errors.New("not found msgid")
}
