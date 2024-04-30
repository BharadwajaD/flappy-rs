package main

import (
	"flag"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/bharadwajaD/flappy-go/pkg/tcp"
	"github.com/rs/zerolog"
)

func main(){

    zerolog.SetGlobalLevel(zerolog.InfoLevel)

    host := flag.String("host", "127.0.0.1" ,"ip address of game engine")
    port := flag.String("port", "42069", "port address of game engine")
    flag.Parse()

    gameOpts := game.NewGameOpts(5, 80, 24);
    serverConfig := tcp.Config{
    	Host:  *host,
    	Port: *port,
    }
    server := tcp.NewServer(&serverConfig)
    server.Run(&gameOpts)
}
