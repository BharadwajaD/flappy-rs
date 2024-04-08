package game

import (
	"log"
	"time"
)

type GameOpts struct {
	frame_rate int
	ticker     *time.Ticker
	win_width  int
	win_height int
}

func NewGameOpts(fps, win_width, win_height int) GameOpts {
	ticker := time.NewTicker(time.Duration(1000/fps) * time.Millisecond)
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
	bird        Bird_t
	pipe        Pipe_t // might be a list
}

func NewGame(opts GameOpts) Game {

	GameOutChan := make(chan Message)
	//GameOutChan <- Message{Obj: Start, Param1: opts.win_width, Param2: opts.win_height}

	return Game{
		GameOutChan: GameOutChan,
		GameInChan:  make(chan Message),
		Stats:       "",
		gameOpts:    opts,
		bird:        NewBird(),
	}
}

func (g *Game) Start() error {

	ticker := g.gameOpts.ticker

	go func() {
		isKeyPressed := false
		for {
			select {
			case inMsg := <-g.GameInChan:
				{
					//TODO: Gamin
					//isKeyPressed = true
					log.Printf("Recieved %v from client\n", inMsg)
				}
			case c := <-ticker.C:
				{
					g.bird.UpdatePos(isKeyPressed)
					g.GameOutChan <- Message{Obj: Bird, Param1: g.bird.xloc, Param2: g.bird.yloc}
					g.GameOutChan <- Message{Obj: Pipe, Param1: c.Second(), Param2: 2}
					isKeyPressed = false
				}
			}

		}
	}()

	return nil
}

func (g *Game) Stop() {
	g.gameOpts.ticker.Stop()
}
