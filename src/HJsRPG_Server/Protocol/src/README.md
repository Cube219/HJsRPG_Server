# Protocol

네트워크 통신에 필요한 프로토콜을 정의합니다.  
(Define a protocol that is needed network communication.)

## BaseProtocol.fbs

기본이 되는 프로토콜입니다.  
(A base protocol.)

각 프로토콜들의 Type을 지정하는 등의 역할을 합니다.  
(Have a role to define a type of each protocols, etc.)

### enum ProtocolType

  * Protocol들의 종류입니다.  
  (The type of protocols.)

  * 각 Protocol(.fbs 파일)들은 이 enum에 대응됩니다.  
  (Each protocols(.fbs file) correspond this enum.)

---

## Client-GameServer

클라이언트와 게임 서버간에 정의된 프로토콜들입니다.  
(Protocols defined between client and game server.)

---

## Client-LoginServer

클라이언트와 로그인 서버간에 정의된 프로토콜들입니다.  
(Protocols defined between client and login server.)

---

## LoginServer-GameServer

로그인 서버와 게임 서버간에 정의된 프로토콜들입니다.  
(Protocols defined between login server and game server.)

---

## ManagementServer-GameServer

관리 서버와 게임 서버간에 정의된 프로토콜들입니다.  
(Protocols defined between management server and game server.)

차후 구현 예정  
(Add in future...)

---

## ManagementServer-LoginServer

관리 서버와 로그인 서버간에 정의된 프로토콜들입니다.  
(Protocols defined between management server and login server.)

차후 구현 예정  
(Add in future...)

---
