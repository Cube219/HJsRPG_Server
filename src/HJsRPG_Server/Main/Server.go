package main

import (
    gs "HJsRPG_Server/GameServer"
)

func main(){
    gameService := gs.CreateGameService(8888)
    gameService.Start()
}