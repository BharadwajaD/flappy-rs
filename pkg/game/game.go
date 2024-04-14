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
		win_width:  win_width,
		win_height: win_height,
	}
}

type Game struct {
	GameOutChan chan Message
	GameInChan  chan Message
	Stats       string
	gameOpts    GameOpts
	bird        Bird_t
	pipe        chan Pipe_t // might be a list
}

func NewGame(opts GameOpts) Game {

	GameOutChan := make(chan Message)

	return Game{
		GameOutChan: GameOutChan,
		GameInChan:  make(chan Message),
		Stats:       "",
		gameOpts:    opts,
		bird:        NewBird(&opts),
		pipe:        GenPipes(&opts),
	}
}

func (g *Game) Start() error {

	ticker := g.gameOpts.ticker
	//TODO: Deal with dis later
	//g.GameOutChan <- Message{Obj: Start, Param1: g.gameOpts.win_width, Param2: g.gameOpts.win_height}

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
			case <-ticker.C:
				{
					g.bird.UpdatePos(isKeyPressed)
                    pipe := <- g.pipe
					g.GameOutChan <- Message{cmd: Bird, params: []int{g.bird.xloc, g.bird.yloc}}
					g.GameOutChan <- Message{cmd: Pipe, params: []int{pipe.xloc, pipe.height}}
					isKeyPressed = false
                    if g.IsCollided(&pipe, &g.bird) {
					    g.GameOutChan <- Message{cmd: End, params: []int{20}}
                    }
				}
			}

		}
	}()

	return nil
}

func (g *Game) Stop() {
	g.gameOpts.ticker.Stop()
}
