package main

import (
	pb "example"
	"fmt"
)

func LoginHandler(msgid uint16, msg interface{}) {
	p := msg.(*pb.RequestLogon)
	fmt.Println("message handler msgid:", msgid, " body:", p)
}

func CreateRoomHandler(msgid uint16, msg interface{}) {
	p := msg.(*pb.MsgTest2)
	fmt.Println("message handler msgid:", msgid, " body:", p)
}

func RegistMsg() {
	RegisterMessage(uint16(pb.MSGID_Logon_Request), &pb.RequestLogon{}, LoginHandler)
	RegisterMessage(uint16(pb.MSGID_CreateRoom_Request), &pb.MsgTest2{}, CreateRoomHandler)
}
