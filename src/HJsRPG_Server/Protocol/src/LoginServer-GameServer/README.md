# LoginServer-GameServer

로그인 서버와 게임 서버간에 정의된 프로토콜들입니다.  
(Protocols defined between login server and game server.)

## SessionInfoProtocol.fbs

세션 정보와 관련된 프로토콜입니다.  
(A protocol related in a session information.)

게임 서버에 세션을 전달하는 역할을 합니다.  
(Have a role to transfer the session to game server.)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---