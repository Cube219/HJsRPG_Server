package GameServer

import (
    "os"
    "time"
    "fmt"
)

// --------------------------
// GameService : 게임 세계의 연산을 담당하는 게임 서비스이다.
type GameService struct {
    ID int64
    logOutput *os.File
}

// Start : 게임 서비스를 실행한다.
func (g *GameService) Start() {
    
    // Log를 남길 파일 생성
    now := time.Now().Format("2006-01-02 15h 04m 05s")
    fmt.Println(now)
    o, err := os.Create("Log(" + now + ").log")
    if err != nil {
        panic(err)
    }
    g.logOutput = o

   g.WriteLog(fmt.Sprintf("GameService(%d) Start.", g.ID))
}

// WriteLog : 로그파일에 로그를 남긴다.
func (g *GameService) WriteLog(log string) {
    now := time.Now().Format("2006-01-02 15h 04m 05s")

    _, err := g.logOutput.WriteString(now + ": " + log + "\n");
    if err != nil {
        panic(err)
    }
}
// ---------------------------

// CreateGameService: 게임 서비스를 생성한다.
func CreateGameService() *GameService {
    g := GameService {
        ID:123,
    }

    return &g
}