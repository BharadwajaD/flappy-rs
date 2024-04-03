package tcp

import (
	"fmt"
	"log"
	"net"

	"github.com/bharadwajaD/flappy-go/pkg/game"
)

// Server ...
type Server struct {
	host string
	port string
}

// Client ...
type Client struct {
	conn net.Conn
	game game.Game
}

// Config ...
type Config struct {
	Host string
	Port string
}

func NewServer(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
			game: game.NewGame(game.NewGameOpts(1000)),
		}
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	//reader := bufio.NewReader(client.conn)
	log.Printf("%v connected\n", client.conn)

    go client.game.Start()

	for gmsg := range client.game.GameChan {
        log.Printf("GameMessage:%v\n", gmsg)
		gmsg, err := game.MessageToStr(&gmsg)
        log.Printf("GameMessage:%v\n", gmsg)
		if err != nil {
			log.Fatalf(err.Error())
		}
		client.conn.Write([]byte(gmsg))
	}
}
