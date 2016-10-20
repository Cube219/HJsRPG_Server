package main

import (
	"log"
	"net"
	"strconv"
	"strings"

	"protocol/users"
	flatbuffers "github.com/google/flatbuffers/go"
)

func ConnectToServer(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// --- 데이터 보냄

/*
	p := &protocol.Test{}
	p.TestString = "test_string"
	p.TestInt = 132

	out, err := proto.Marshal(p)
	if err != nil{
		log.Fatalf("Fail to encode.\n %s", err)
	}*/

	b := flatbuffers.NewBuilder(0)
	b.Reset();

	name_position := b.CreateByteString([]byte("NName"))

	users.UserStart(b)
	users.UserAddName(b, name_position) 
	users.UserAddId(b, 132)
	user_pos := users.UserEnd(b)

	b.Finish(user_pos) 
	
	raw := b.Bytes[b.Head():]

	conn.Write(raw)
}

func main(){
	ConnectToServer("127.0.0.1", 8888)
}