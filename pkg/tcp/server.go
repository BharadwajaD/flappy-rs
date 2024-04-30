package tcp

import (
	"fmt"
	"net"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/rs/zerolog/log"
)

//multiple clients

type Server struct {
	host string
	port string
}

func NewServer(config *Config) Server {
	return Server{
		host: config.Host,
		port: config.Port,
	}
}

func (s *Server) Run(gameOpts *game.GameOpts) {

	ggame := game.NewGroupGame(gameOpts)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))

	if err != nil {
		log.Fatal().Msgf("%v\n", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal().Msgf("%v\n", err)
		}

		log.Debug().Msgf("DEBUG:SERVER RUN: %+v", conn)
		newgame := game.NewGame(&ggame, gameOpts.Clone())
		log.Debug().Msgf("DEBUG:SERVER RUN: %+v", newgame)
		client := &Client{
			conn: conn,
			game: newgame,
		}
		go client.handleRequest()
	}
}
