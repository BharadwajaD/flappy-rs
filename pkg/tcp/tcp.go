package tcp

import (
	"bufio"
	"net"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/rs/zerolog/log"
)


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
