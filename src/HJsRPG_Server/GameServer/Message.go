package GameServer

import "encoding/binary"

// Message : 프로토콜에 있는 메시지입니다.
type Message struct {
    protocolType uint
    messageType uint
    data []byte
}

// ToByte : 메시지를 바이트 배열로 바꿉니다.
func (m *Message) ToByte() []byte {
    // uint들을 byte배열로 바꿔줌
    protocolTypeByte := make([]byte, 4)
    binary.LittleEndian.PutUint32(protocolTypeByte, uint32(m.protocolType)) 
    messageTypeByte := make([]byte, 4)
    binary.LittleEndian.PutUint32(messageTypeByte, uint32(m.messageType)) 

    // protocolType, messageType 그리고 data를 하나로 합침
    d := make([]byte, 8 + len(m.data))
    d = append(d, protocolTypeByte...)
    d = append(d, messageTypeByte...)
    d = append(d, m.data...)

    return d
}