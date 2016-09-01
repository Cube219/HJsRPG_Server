package main

import (
    "net"
    "log"
    "strconv"
    "bufio"
    "io"
)

// 소켓 서버를 연다
func OpenSocketServer(port int) {
    listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
    defer listen.Close()

    if err != nil {
        log.Fatalf("Fail to open socket in port %d\n%s", port, err)
        return
    }

    log.Printf("Success to open socket in port %d", port)

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

    var (
        buf = make([]byte, 1024)
        r = bufio.NewReader(conn)
    )

LOOP:
    for {
        n, err := r.read(buf)
        data := string(buf[:n])

        switch err {
            case io.EOF:
                break LOOP
            
            case nil:
                log.Println("Received:", data)
            default:
                log.Fatalf("Fail to receive data\n%s", err)
                return
        }
    }
}

func main(){
    OpenSocketServer(8888)
}