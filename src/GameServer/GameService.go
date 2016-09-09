package GameServer

// --------------------------
// GameService: 게임 세계의 연산을 담당하는 게임 서비스이다.
type GameService struct {
    ID int64
}

// GameService 관련 함수

// ---------------------------



// CreateGameService: 게임 서비스를 생성한다.
func CreateGameService() *GameService {
    g := GameService {
        ID:123
    }

    return &g
}