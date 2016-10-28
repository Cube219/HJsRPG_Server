package main

import (
    gs "GameServer"
)

func main(){
    gameService := gs.CreateGameService()
    gameService.Start()
}