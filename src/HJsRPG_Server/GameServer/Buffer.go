package GameServer

import (
    flatbuffers "github.com/google/flatbuffers/go"
)

// ---------------------------------------------------------

// ProtocolType : 어떤 프로토콜인지 구분하는 역할
type ProtocolType uint
const (
    ProtocolType_TestProtocol ProtocolType = iota
)

// Protocol : 클라이언트 / 서버간 통신을 할 내용
type Protocol struct {
    protocolType uint
    protocol flatbuffers.Table
}

// ---------------------------------------------------------

// Buffer : 클라이언트 / 서버간 통신을 할 버퍼
type Buffer struct {
    protocols []Protocol
}

// Flush : 버퍼에 쓴 내용을 byte로 바꿈
func (b *Buffer) Flush() []byte {
    // Flatbuffer 빌더 생성
    builder := flatbuffers.NewBuilder(0) 
    builder.Reset()

  //  builder.fini

    for _, proto := range b.protocols {

    }
    return nil;
}

// ---------------------------------------------------------

func a (){
}