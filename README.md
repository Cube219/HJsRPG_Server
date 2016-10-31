# HJsRPG_Server

Golang을 공부하기 위한 RPG서버 입니다.  
(RPG Server to study Golang.)

## 실행 방법 (How to Run)

## 구조 (Structure)

* **HJsRPG_Server**
  * **Main**
  * **GameServer**
    * **GameService.go**
    * **World.go**
    * **Map.go**
    * **GameDB.go**
  * **LoginServer**
    * **LoginService.go**
    * **AccountDB.go**
  * **ManagementServer**
    * **ManagementService.go**
  * **GameObject**
    * **BaseObject.go**
    * **Actor.go** : **BaseObject.go**
    * **Charactor.go** : **Actor.go**
    * **Monster.go** : **Actor.go**
    * **Item.go** : **BaseObject.go**
  * **Network**
    * **ConnectedClient.go**
    * **Session.go**
    * **Message.go**
  * **Protocol**

더 자세한 내용은 Document/ProgramStructure.md 파일을 참조해 주세요.  
(For more information, check Document/ProgramStructure.md file.)

## 사용한 라이브러리 (Using Library)

* Flatbuffers (<https://google.github.io/flatbuffers/>
* golang/crypto (<https://github.com/golang/crypto>)

## 참고한 사이트 (Reference)

* kasworld/go4game (<https://github.com/kasworld/go4game>)