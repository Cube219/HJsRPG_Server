# InGame

실제 게임 내부와 관련된 프로토콜들입니다.  
(Protocols related in actual game.)

## InGameBaseProtocol.fbs

InGame에서 기본이 되는 프로토콜입니다.  
(A base protocol from InGame.)

각종 struct나 enum들의 정의로 구성되어 있습니다.  
(It is composed of the definition of structs or enums)

---

## MapActionProtocol.fbs

맵에 있는 것들의 활동과 관련된 프로토콜입니다.  
(A protocol related to action of the things in a map.)

몬스터의 움직임 / 다른 캐릭터의 이동...  
(Movement of monster, of others characters...)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---

## MapInfoProtocol.fbs

맵에 있는 것들의 정보와 관련된 프로토콜입니다.  
(A protocol related to information of the things in a map.)

몬스터 / 다른 캐릭터...  
(Monster / other characters...)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---

## PlayerActionProtocol.fbs

플레이어의 활동과 관련된 프로토콜입니다.  
(A protocol related to action of the player.)

플레이어의 이동 / 스킬 사용 / 아이템 습득...  
(Movement of player / using a skill / get a item...)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---

## PlayerInfoProtocol.fbs

플레이어의 정보와 관련된 프로토콜입니다.  
(A protocol related to information of the player.)

스킬 / 가지고 있는 아이템 / 스텟...  
(Skill / belonging item / state...)

### enum MessageType

  * 이 프로토콜 안에 있는 Message들의 종류입니다.  
  (The type of messages in this protocol.)

  * 각 Message들은 이 enum에 대응됩니다.  
  (Each messages correspond this enum.)

---

