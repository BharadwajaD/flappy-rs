package main

import (
	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/bharadwajaD/flappy-go/pkg/tcp"
)

func main(){
    gameOpts := game.NewGameOpts(1, 80, 24);
    serverConfig := tcp.Config{
    	Host: "127.0.0.1",
    	Port: "42069",
    }
    server := tcp.NewServer(&serverConfig)
    server.Run(&gameOpts)
}
