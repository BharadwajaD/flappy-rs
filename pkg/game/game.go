package game

import (
	//	"log"
	"log"
	"time"
)

type GameOpts struct {
	frame_rate int
	ticker     *time.Ticker
}

func NewGameOpts(fps int) GameOpts {
	ticker := time.NewTicker(time.Duration(fps) * time.Millisecond)
	return GameOpts{
		frame_rate: fps,
		ticker:     ticker,
	}
}

type Game struct {
	GameOutChan chan Message
	GameInChan  chan Message
	Stats       string
	gameOpts    GameOpts
}

func NewGame(opts GameOpts) Game {
	return Game{
		GameOutChan: make(chan Message),
		GameInChan:  make(chan Message),
		Stats:       "",
		gameOpts:    opts,
	}
}

func (g *Game) Start() error {

	ticker := g.gameOpts.ticker

	go func() {
		for {
			select {
            case inMsg := <-g.GameInChan:
				{
                    //TODO: Gaming
					log.Printf("Recieved %v from client\n", inMsg)
				}
			case c := <-ticker.C:
				{
					g.GameOutChan <- Message{obj: Pipe, param1: c.Second(), param2: 2}
					g.GameOutChan <- Message{obj: Bird, param1: c.Second(), param2: 18}
				}
			}

		}
	}()

	return nil
}

func (g *Game) Stop() {
	g.gameOpts.ticker.Stop()
}
