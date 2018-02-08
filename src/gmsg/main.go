package main

import (
	pb "example"

	"github.com/golang/protobuf/proto"
)

func main() {
	RegistMsg()

	t1 := pb.RequestLogon{
		Id:    11111,
		Name:  "name1111",
		Email: "1111@example.com",
	}
	out, _ := proto.Marshal(&t1)
	HandleRawData(uint16(pb.MSGID_Logon_Request), out)

	t2 := pb.MsgTest2{
		Id:    22222,
		Name:  "name222",
		Email: "2222@example.com",
	}
	out, _ = proto.Marshal(&t2)
	HandleRawData(uint16(pb.MSGID_CreateRoom_Request), out)
}
