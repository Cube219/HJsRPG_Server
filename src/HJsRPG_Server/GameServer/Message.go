package GameServer

import "encoding/binary"

// Message : 프로토콜에 있는 메시지입니다.
type Message struct {
    ProtocolType uint
    MessageType uint
    Data []byte
}

// ToByte : 메시지를 바이트 배열로 바꿉니다.
func (m *Message) ToByte() []byte {
    // uint들을 byte배열로 바꿔줌
    protocolTypeByte := make([]byte, 4)
    binary.LittleEndian.PutUint32(protocolTypeByte, uint32(m.ProtocolType)) 
    messageTypeByte := make([]byte, 4)
    binary.LittleEndian.PutUint32(messageTypeByte, uint32(m.MessageType)) 

    // protocolType, messageType 그리고 data를 하나로 합침
    d := make([]byte, 8 + len(m.Data))
    d = append(d, protocolTypeByte...)
    d = append(d, messageTypeByte...)
    d = append(d, m.Data...)

    return d
}

// Parse : 바이트 배열을 메시지로 바꿔준다.
func Parse(d []byte) Message {
    var m Message

    protocolTypeByte := d[0:4]
    messageTypeByte := d[4:8]
    data := d[8:]

    m.ProtocolType = uint(binary.LittleEndian.Uint32(protocolTypeByte))
    m.MessageType = uint(binary.LittleEndian.Uint32(messageTypeByte))
    m.Data = data

    return m
}