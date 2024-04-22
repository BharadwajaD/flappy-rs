package tcp

import (
	"bufio"
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

func (server *Server) Run(gameOpts *game.GameOpts) {
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
			game: game.NewGame(gameOpts.Clone()),
		}
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	log.Printf("%v connected\n", client.conn)

    go client.game.Start()
    go func(){
        for gmsg := range client.game.GameOutChan {
            gmsg, err := game.MessageToStr(&gmsg)
            if err != nil {
                log.Fatalf(err.Error())
            }
            log.Printf("Sent: %s\n", gmsg)
            client.conn.Write([]byte(gmsg))
        }
    }()

    reader := bufio.NewReader(client.conn)
    for {
        msg, err := reader.ReadString('?')
        if err != nil {return}
        gmsg , err := game.MessageFromStr(string(msg))
        if err != nil {return}
        log.Printf("DEBUG:HANDLE REQUEST: %+v\n", gmsg )
        client.game.GameInChan <- *gmsg
    }
}
