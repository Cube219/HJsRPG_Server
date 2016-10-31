# 프로그램 구조 (Program Structure)

작성/번역중...  
(writing/translating...)

* **HJsRPG_Server**
  * **Main**
    * 서버를 처음 실행하면 실행되는 곳.
  * **GameServer**
    * **GameService.go**
      * 실제 게임 서버의 코어 부분.
      * World간 이동을 관여하거나 세션을 관리한다.
    * **World.go**
      * 말 그대로 게임세계이다.
      * Map간의 이동을 관여하거나 채팅을 관리한다.
    * **Map.go**
      * World안에 있는 맵들이다.
      * 주요 판정 체크 / 몬스터 관리 / 확률 등을 관리한다.
    * **GameDB.go**
      * 게임DB하고 통신한다.
  * **LoginServer**
    * **LoginService.go**
      * 맨 처음에 접속할 때 로그인하는 서버이다.
      * 아이디 / 비밀번호 정보를 확인하고, 세션을 생성해준다.
      * 그리고 만든 세션을 GameServer에 전송한다.
    * **AccountDB.go**
      * 계정과 관련된 DB하고 통신한다.
  * **ManagementServer**
    * **ManagementService.go**
      * 서버를 관리하는 서버이다.
      * 서버상태를 관리하거나 표시한다.
  * **GameObject**
    * **BaseObject.go**
      * 게임내에 존재하는 기본적인 오브젝트이다.
    * **Actor.go** : **BaseObject.go**
      * 게임내에 돌아다니는 것들이다.
    * **Charactor.go** : **Actor.go**
      * 캐릭터이다.
    * **Monster.go** : **Actor.go**
      * 몬스터이다.
    * **Item.go** : **BaseObject.go**
      * 게임내에 있는 아이템이다.
  * **Network**
    * **ConnectedClient.go**
      * 연결된 클라이언트 정보이다.
    * **Session.go**
      * 클라이언트에 부여된 세션이다.
    * **Message.go**
      * 서로간에 통신할 메시지이다.
      * 통신 버퍼라고 생각해도 됨.
      * Protocol을 []byte로 변환/복원, 암호화를 담당한다.
  * **Protocol**
    * 서버간 / 서버와 클라이언트간 통신하는 프로토콜을 정의한다.
    * Flatbuffers에 의해 내용이 자동으로 생성됨.


차후 구조간의 상관관계 추가예정...  
(Add relationship between each structures in future...)
