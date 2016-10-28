package GameServer 

import (
    "net"
)

// ---------------------------------------------------------

// ConnectedClient : 연결된 클라이언트이다.
type ConnectedClient struct {
    readCh chan Message
    sendCh chan Message
    
    conn net.Conn
}

// Send : 해당 클라이언트로 보낸다.
func (c *ConnectedClient) Send(message Message){
    // 보냄
    c.conn.Write(message.ToByte()) 
}

// ReadLoop : 들어오는 정보를 받아내는 루프문
func (c *ConnectedClient) ReadLoop() {
    defer close(c.readCh)
}

// SendLoop : 클라이언트로 정보를 보내는 루프문
func (c *ConnectedClient) SendLoop() {
    defer close(c.sendCh);

    for {
        sendMessage := <- c.sendCh
        c.Send(sendMessage) 
    }
}

// ---------------------------------------------------------

// NewConnectedClient : 클라이언트 연결에 관한 새로운 정보 생성
func NewConnectedClient(conn net.Conn) *ConnectedClient {
    c := ConnectedClient {
        conn: conn,
        readCh: make(chan Message, 1),
        sendCh: make(chan Message, 1),
    }

    go c.ReadLoop();
    go c.SendLoop();

    return &c;
}