package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/rs/zerolog/log"
)

type TCPServer struct {
	host string
	port string
}

type tcprw struct {
	rw *bufio.ReadWriter
}


var stream string //TODO: SHould not use global var

func (trw tcprw) ReadString(delim byte) (string, error) {
    chunk, err := trw.rw.Reader.ReadString(delim) 
    stream += chunk
    if err != nil {
        //no delim in the read data
        return "", err
    }

	idx := strings.Index(stream, "?")
	cmd := stream[:idx]
	stream = stream[idx+1:]

	return cmd, err
}

func (trw tcprw) WriteString(str string) error {
    _, err := trw.rw.Writer.WriteString(str)
    err = trw.rw.Flush()

    return err
}

func NewTCPServer(config *Config) TCPServer {
	return TCPServer{
		host: config.Host,
		port: config.Port,
	}
}

func (s TCPServer) Run(gameOpts *game.GameOpts) {

	ggame := game.NewGroupGame(gameOpts)
	client_id := 0
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

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)

		client_id++
		client := &Client{
			id: "TCP:" + strconv.Itoa(client_id),
			rw: tcprw{rw: bufio.NewReadWriter(reader, writer)},

			conn: conn,
			game: newgame,
		}

		go client.handleRequest()
	}
}
