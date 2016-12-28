# 프로그램 구조 (Program Structure)

* **HJsRPG_Server**
  * **Main**
    * 서버를 처음 실행하면 실행되는 곳.  
      (Run at first when you run a server.)
  * **GameServer**
    * **GameService.go**
      * 실제 게임 서버의 코어 부분.  
        (A core in game server.)
      * 플레이어의 World간 이동을 관여하거나 세션을 관리한다.  
        (Control player's movement between Worlds or session.)
    * **World.go**
      * 말 그대로 게임세계이다.  
        (Just a game world.)
      * 플레이어의 Map간의 이동을 관여하거나 채팅을 관리한다.  
        (Control player's movement between Maps or chat.)
    * **Map.go**
      * World안에 있는 맵이다.  
        (A map in the World.)
      * 충돌 체크 / 몬스터 / 확률 등을 관리한다.  
        (Control collision check / monster / probability.)
    * **GameDB.go**
      * 게임DB하고 통신한다.  
        (Communicate with GameDB.)
  * **LoginServer**
    * **LoginService.go**
      * 로그인하는 서버이다.  
        (A server that control login.)
      * 아이디 / 비밀번호 정보를 확인하고, 세션을 생성해준다.  
        (Check ID / password, and create session.)
      * 그리고 만든 세션을 GameService에 전송한다.  
        (Then transfer the session to GameService.)
    * **AccountDB.go**
      * 계정과 관련된 DB하고 통신한다.  
        (Communicate with DB related account.)
  * **ManagementServer**
    * **ManagementService.go**
      * 게임서버를 관리하는 서버이다.  
        (A server that manage the game server.)
      * 서버상태를 관리하거나 표시한다.  
        (Show and manage server states.)
  * **GameObject**
    * **BaseObject.go**
      * 게임내에 존재하는 기본적인 오브젝트이다.  
        (A base object that exists in the game.)
    * **Actor.go** : **BaseObject.go**
      * 게임내에서 움직일 수 있는 것들이다.  
        (Something that can move in the game.)
    * **Charactor.go** : **Actor.go**
      * 캐릭터이다.  
        (A character.)
    * **Monster.go** : **Actor.go**
      * 몬스터이다.  
        (A monster.)
    * **Item.go** : **BaseObject.go**
      * 게임내에 있는 아이템이다.  
        (A item in the game.)
  * **Network**
    * **ConnectedClient.go**
      * 연결된 클라이언트 정보이다.  
        (A information about connected client.)
    * **Session.go**
      * 클라이언트에 부여된 세션이다.  
        (A session given to client.)
    * **Message.go**
      * 서로간에 통신할 메시지이다.  
        (A message used to communicate each others.)
      * 통신 버퍼라고 생각해도 됨.  
        (Just like communication buffer.)
      * Protocol을 []byte로 변환/복원, 암호화를 한다.  
        (Convert between Protocol and []byte data, and encrypt it.)
  * **Protocol**
    * 서버간 / 서버와 클라이언트간 통신하는 프로토콜을 정의한다.  
      (Define protocols between each servers / server and client.)
    * Flatbuffers에 의해 내용이 자동으로 생성됨.  
      (Created automatically by Flatbuffers.)
    * **BaseProtocol.fbs**
      * 프로토콜의 베이스 기반  
        (A base protocol.)
      * 각 프로토콜들의 enum지정을 한다.  
        (Assign enums each protocols.)
    * **Client-LoginServer**
      * 클라이언트-로그인서버 관련 프로토콜.  
        (A protocol between Client and LoginServer.)
      * **AssignProtocol.fbs**
        * 회원가입과 관련된 프로토콜  
          (A protocol related assignment.)
      * **LoginProtocol.fbs**
        * 로그인과 관련된 프로토콜  
          (A protocol related login.)
        * 아이디 / 비밀번호, 세션...  
          (ID / password, session...)
    * **Client-GameServer**
      * 클라이언트-게임서버 관련 프로토콜.  
        (A protocol between Client and GameServer.)
      * **InGame**
        * **InGameBaseProtocol.fbs**
          * InGame에 있는 프로토콜의 기본 프로토콜.  
            (A protocol that base protocols in InGame.)
          * 각종 struct나 enum들의 정의들이 있음.  
            (Has various definition of struct and enum.)
        * **PlayerActionProtocol.fbs**
          * 플레이어 홛동과 관련된 프로토콜.  
            (A protocol related player's activity.)
          * 플레이어 이동 / 스킬사용 / 아이템 습득 등등...  
            (Player's movement / use skills, get items...)
        * **PlayerInfoProtocol.fbs**
          * 플레이어 정보과 관련된 프로토콜.  
            (A protocol related the information of player.)
          * 아이템 / 스킬 / 스텟...  
            (Item / skill / state...)
        * **MapActionProtocol.fbs**
          * 맵에 있는 것들의 활동과 관련된 프로토콜.  
            (A protocol related activity of something in the map.)
          * 몬스터 움직임 / 다른 캐릭터 이동...  
            (Monster's movement / other characters' movement...)
        * **MapInfoProtocol.fbs**
          * 맵에 있는 정보과 관련된 프로토콜.  
            (A protocol related information in the map.)
          * 몬스터 / 다른 캐릭터...  
            (Monster / other characters...)
      * **Chat** (차후 구현)(In future)
        * **MapChatProtocol.fbs**
          * 일반 Map에서 채팅할 때 쓰는 프로토콜.  
            (A protocol used to chat in Map.)
        * **WorldChatProtocol.fbs**
          * World 내에서 채팅할 때 쓰는 프로토콜.  
            (A protocol used to chat in World.)
          * 해당 World에 속해있는 모두에게 보여짐.  
            (Show everyone in the World.)
        * **WideChatProtocol.fbs**
          * 게임에 있는 모두에게 채팅할 때 쓰는 프로토콜.  
            (A protocol used to chat everyone in the game.)
    * **LoginServer-GameServer**
      * 로그인서버-게임서버 관련 프로토콜.  
        (A protocol between LoginServer and GameServer.)
      * **SessionInfoProtocol.fbs**
        * 세션 관련 정보를 보낼 때 쓰는 프로토콜.  
          (A protocol used to send session information.)
        * 보통 로그인하고 나서 보냄.  
          (Usually used after login.)
    * **ManagementServer-LoginServer**
      * 관리서버-로그인서버 관련 프로토콜. (차후 구현)  
        (A protocol between ManagementServer and LoginServer)(In future)
    * **ManagementServer-GameServer**
      * 관리서버-게임서버 관련 프로토콜. (차후 구현)  
        (A protocol between ManagementServer and GameServer)(In future)

차후 구조간의 상관관계 추가예정...  
(Add relationship between each structures in future...)
