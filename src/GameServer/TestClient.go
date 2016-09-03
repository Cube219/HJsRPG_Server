package main

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	protocol "Protocol"
)

func ConnectToServer(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// --- 데이터 보냄

	p := &protocol.Test{}
	p.TestString = "test_string"
	p.TestInt = 132

	out, err := proto.Marshal(p)
	if err != nil{
		log.Fatalf("Fail to encode.\n %s", err)
	}

	conn.Write(out)
}

func main(){
	ConnectToServer("127.0.0.1", 8888)
}