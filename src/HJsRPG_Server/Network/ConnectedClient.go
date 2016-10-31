package Network

import (
    "net"
	"bufio"
	"io"
	"fmt"
    baseProtocol "HJsRPG_Server/Protocol/GameServer"
    testProtocol "HJsRPG_Server/Protocol/GameServer/TestProtocol"
)

// ---------------------------------------------------------

// ConnectedClient : 연결된 클라이언트이다.
type ConnectedClient struct {
    ReadCh chan Message
    SendCh chan Message
    
    conn net.Conn
}

// Send : 해당 클라이언트로 보낸다.
func (c *ConnectedClient) Send(message Message){
    // 보냄
    c.conn.Write(message.ToByte()) 
}

// SendLoop : 클라이언트로 정보를 보내는 루프문
func (c *ConnectedClient) SendLoop() {
    defer close(c.SendCh);

    for {
        sendMessage := <- c.SendCh
        c.Send(sendMessage) 
    }
}

// ReadLoop : 들어오는 정보를 받아내는 루프문
func (c *ConnectedClient) ReadLoop() {
    defer close(c.ReadCh)

    var (
        buf = make([]byte, 1024)
        r = bufio.NewReader(c.conn) 
    )

READLOOP:
    for {
        n, err := r.Read(buf)
        data := buf[:n]

        switch err {
            case io.EOF:
                break READLOOP

            case nil:
                c.ProcessMessage(data)

            default:
        }
    }
}

// ProcessMessage : 메시지를 처리한다
func (c *ConnectedClient) ProcessMessage(data []byte) {
    msg := Parse(data) 
    fmt.Printf("%d %d\n", msg.ProtocolType, msg.MessageType)
    // 메시지에 따라서 처리 부분 구현

    switch msg.ProtocolType {
        case baseProtocol.ProtocolTypeTestProtocol: // TestProtocol
            processInTestProtocol(msg) 
    }
}

// processInTestProtocol : TestProtocol에 있는 내용을 처리한다.
func processInTestProtocol(msg Message) {
    switch msg.MessageType {
        case testProtocol.MessageTypeTestPing: // TestPing
            testPingMsg := testProtocol.GetRootAsTestPing(msg.Data, 0) 
            fmt.Printf("Get 'Ping' message in TestProtocol\n")
            fmt.Printf("Value : %d\n", testPingMsg.IntValue())  
    }
}

// ---------------------------------------------------------

// NewConnectedClient : 클라이언트 연결에 관한 새로운 정보 생성
func NewConnectedClient(conn net.Conn) *ConnectedClient {
    c := ConnectedClient {
        conn: conn,
        ReadCh: make(chan Message, 1),
        SendCh: make(chan Message, 1),
    }

    go c.ReadLoop();
    go c.SendLoop();

    return &c;
}