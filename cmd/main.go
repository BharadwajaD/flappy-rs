package main

import "github.com/bharadwajaD/flappy-go/pkg/tcp"

func main(){
    server := tcp.NewServer(&tcp.Config{
    	Host: "127.0.0.1",
    	Port: "8080",
    })

    server.Run()
}
