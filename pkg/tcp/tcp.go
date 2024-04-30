package tcp

import (
	"bufio"
	"fmt"
	"net"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/rs/zerolog/log"
)

// Server_t ...
type Server_t struct {
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

func NewServer_t(config *Config) *Server_t {
	return &Server_t{
		host: config.Host,
		port: config.Port,
	}
}

func (server *Server_t) Run(gameOpts *game.GameOpts) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal().Msgf("%+v\n", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal().Msgf("%+v\n", err)
		}

		client := &Client{
			conn: conn,
			game: game.NewGame(nil, gameOpts.Clone()),
		}
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	log.Info().Msgf("%v connected\n", client.conn)

	go client.game.Start()
	go func() {
		for gmsg := range client.game.GameOutChan {
			gmsg, err := game.MessageToStr(&gmsg)
			if err != nil {
				log.Fatal().Msgf(err.Error())
			}
			log.Info().Msgf("Sent: %s\n", gmsg)
			client.conn.Write([]byte(gmsg))
		}
	}()

	reader := bufio.NewReader(client.conn)
	for {
		msg, err := reader.ReadString('?')
		if err != nil {
			return
		}
		gmsg, err := game.MessageFromStr(string(msg))
		if err != nil {
			return
		}
		log.Debug().Msgf("DEBUG:HANDLE REQUEST: %+v\n", gmsg)
		client.game.GameInChan <- *gmsg
	}
}
