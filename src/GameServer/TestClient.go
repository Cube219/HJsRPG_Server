package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

func ConnectToServer(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte("Ping"))
}

func main() {
	ConnectToServer("127.0.0.1", 8888)
}
