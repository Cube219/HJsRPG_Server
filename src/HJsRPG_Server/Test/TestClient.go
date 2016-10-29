package main

import (
	"log"
	"net"
	"strconv"
	"strings"

	flatbuffers "github.com/google/flatbuffers/go"
	testProtocol "HJsRPG_Server/Protocol/GameServer/TestProtocol"
)

func ConnectToServer(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	b := flatbuffers.NewBuilder(0)
	b.Reset();

	testProtocol.TestPingStart(b)
	testProtocol.TestPingAddIntValue(b, 15)
	pingPos := testProtocol.TestPingEnd(b)

	b.Finish(pingPos) 
	
	raw := b.Bytes[b.Head():]

	conn.Write(raw)
}

func main(){
	ConnectToServer("127.0.0.1", 8888)
}