package main

import (
    "fmt"
    gs "GameServer"
)

func main(){
    gameService := gs.CreateGameService()
    fmt.Println(gameService.ID)
}