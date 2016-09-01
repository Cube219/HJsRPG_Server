package main

import (
    "net"
    "fmt"
    "strconv"
)

// 소켓 서버를 연다
func OpenSocketServer(port int) {
    listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
    defer listen.Close()

    if err != nil {
        fmt.Println("Fail to open socket in port", port, "\n", err)
        return
    }

    fmt.Println("Success to open socket in port", port)

    for {
        conn, err := listen.Accept()

        if err != nil {
            println(err)
            continue
        }

        go handler(conn)
    }
}

func handler(conn net.Conn){
    defer conn.Close()
}

func main(){
    OpenSocketServer(8888)
}