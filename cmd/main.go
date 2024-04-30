package main

import (

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/bharadwajaD/flappy-go/pkg/tcp"
	"github.com/rs/zerolog"
)

func main(){

    zerolog.SetGlobalLevel(zerolog.InfoLevel)

    gameOpts := game.NewGameOpts(5, 80, 24);
    serverConfig := tcp.Config{
    	Host: "127.0.0.1",
    	Port: "42069",
    }
    server := tcp.NewServer(&serverConfig)
    server.Run(&gameOpts)
}
