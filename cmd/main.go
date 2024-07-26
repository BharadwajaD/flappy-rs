package main

import (
	"flag"
	"log"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/bharadwajaD/flappy-go/pkg/server"
	"github.com/rs/zerolog"
)

func main(){

    zerolog.SetGlobalLevel(zerolog.DebugLevel)

    host := flag.String("host", "127.0.0.1" ,"ip address of game engine")
    port := flag.String("port", "42069", "port address of game engine")
    stype := flag.String("type", "TCP", "type of server")


    flag.Parse()

    gameOpts := game.NewGameOpts(1, 80, 24);
    serverConfig := server.Config{
    	Host:  *host,
    	Port: *port,
    }

    server, err := server.NewGameServer(*stype, &serverConfig)
    if err != nil{
        log.Fatal(err.Error())
    }
    server.Run(&gameOpts)
}
