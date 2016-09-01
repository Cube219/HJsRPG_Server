package main

import (
    "net"
    "fmt"
    "strconv"
)

// 소켓 서버를 연다
func OpenSocketServer(port int)
{
    listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
    defer listen.Close()

    if err != nil {
        println("Fail to open socket in port %d\n%s", port, err)
        return
    }

    println("Success to open socket in port %d", port)

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

func main{
    OpenSocketServer(8888)
}