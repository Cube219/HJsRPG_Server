# Chat

채팅과 관련된 프로토콜들입니다.  
(Protocols related to chat.)

## MapChatProtocol.fbs

맵 상에서 채팅을 할 때 쓰이는 프로토콜입니다.  
(A protocol using chatting in a map.)

이 채팅은 같은 맵에 있는 사람에게만 보여집니다.  
(This chatting is appeared only people in a same map.)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---

## WideChatProtocol.fbs

Wide 채팅을 할 때 쓰이는 프로토콜입니다.  
(A protocol using a wide chatting.)

이 채팅은 게임에 접속한 모든 사람에게 보여집니다.  
(This chatting is appeared all people in the game.)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---

## WorldChatProtocol.fbs

World 상에서 채팅을 할 때 쓰이는 프로토콜입니다.  
(A protocol using chatting in a world.)

이 채팅은 같은 World에 있는 사람에게만 보여집니다.  
(This chatting is appeared only people in a same world.)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---