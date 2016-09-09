package GameServer

import (
  //  "os"
 //   "time"
//    "fmt"
)

// --------------------------
// GameService: 게임 세계의 연산을 담당하는 게임 서비스이다.
type GameService struct {
    ID int64
 //   logOutput *File
}

// GameService 관련 함수
// Start: 게임 서비스를 실행한다.
func (g *GameService) Start() {
    /*
    // Log를 남길 파일 생성
    now := time.Now()
    fo, err := os.Create("Log/Log(" + now.String() + ").log")
    if err != nil {
        panic(err)
    }
    g.logOutput = fo*/
}
/*
// Log: 로그파일에 로그를 남긴다.
func (g *GameService) Log(format string, a ...interface{}) {
    now := time.Now()

    _, err := g.logOutput.WriteString(now.String() + ": " + fmt.Sprintf(format, a));
    if err != nil {
        panic(err)
    }
}*/

// ---------------------------



// CreateGameService: 게임 서비스를 생성한다.
func CreateGameService() *GameService {
    g := GameService {
        ID:123,
    }

    return &g
}