package server

import (
	"net"

	"github.com/bharadwajaD/flappy-go/pkg/game"
	"github.com/rs/zerolog/log"
)

type rwI interface{
    ReadString(byte) (string, error)
    WriteString(string) error
}

// Client ...
type Client struct {
	id string

    rw rwI

	conn net.Conn
	game game.Game
}

func (client *Client) handleRequest() {

	log.Info().Msgf("%v connected\n", client.id)

	go client.game.Start()

	go func() {
		//from game to client
		for gmsg := range client.game.GameOutChan {
			gmsg, err := game.MessageToStr(&gmsg)

			if err != nil {
				log.Error().Msgf("Error :%s\n", err.Error())
				break // break or continue ??
			}
			log.Info().Msgf("DEBUG:CLIENT SENT: %+v\n", gmsg)
			err = client.rw.WriteString(gmsg)

			if err != nil {
				log.Error().Msgf("Error :%s\n", err.Error())
				break
			}
		}
	}()

	for {
		//from client to game
		msg, err := client.rw.ReadString('?') //should read till '?'
		if err != nil {
			return
		}
		gmsg, err := game.MessageFromStr(msg)
        log.Debug().Msgf("DEBUG:CLIENT RECEIVE: %s and %+v\n", msg, gmsg)
        if gmsg == nil {
            continue
        }
		if err != nil {
			return
		}
		log.Debug().Msgf("DEBUG:CLIENT RECEIVE: %+v\n", gmsg)
		client.game.GameInChan <- *gmsg
	}
}

func (client *Client) Close() {
}
