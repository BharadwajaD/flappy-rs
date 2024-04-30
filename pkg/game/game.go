package game

import (
	"time"

	"github.com/rs/zerolog/log"
)

type GameOpts struct {
	frame_rate int
	ticker     *time.Ticker
	win_width  int
	win_height int
}

func NewGameOpts(fps, win_width, win_height int) GameOpts {
	return GameOpts{
		frame_rate: fps,
		ticker:     time.NewTicker(time.Duration(1000/fps) * time.Millisecond),
		win_width:  win_width,
		win_height: win_height,
	}
}

func (gopts *GameOpts) Clone() GameOpts {
	return GameOpts{
		frame_rate: gopts.frame_rate,
		ticker:     time.NewTicker(time.Duration(1000/gopts.frame_rate) * time.Millisecond),
		win_width:  gopts.win_width,
		win_height: gopts.win_height,
	}
}

type Game struct {
	GameOutChan chan Message
	GameInChan  chan Message
	stats       int
	gameOpts    GameOpts
	bird        Bird_t
	pipe        <-chan Pipe_t
	game_id     int
	ggame       *GroupGame
}

var game_id int

func NewGame(ggame *GroupGame, opts GameOpts) Game {

	game_id++
	return Game{
		GameOutChan: make(chan Message),
		GameInChan:  make(chan Message),
		stats:       0,
		gameOpts:    opts,
		bird:        NewBird(&opts),
		pipe:        GenPipes(ggame, &opts, game_id),
		game_id:     game_id,
		ggame:       ggame,
	}
}

type GameStatus int

const (
	Collided GameStatus = iota
	Crossed
	FarAway
)

func (g *Game) Status(b *Bird_t, p *Pipe_t) GameStatus {
	if p.xloc == b.xloc {
		if b.yloc <= p.height || b.yloc >= g.gameOpts.win_height-p.height {
			//return Collided
            return Crossed //For testing 
		} else {
			return Crossed
		}
	}

	return FarAway
}

func (g *Game) Start() error {

	ticker := g.gameOpts.ticker
	//TODO: Deal with dis later
	//g.GameOutChan <- Message{Obj: Start, Param1: g.gameOpts.win_width, Param2: g.gameOpts.win_height}

	go func() {
		isKeyPressed := false
		pipe := <-g.pipe
		//starting pipe and bird
		log.Debug().Msgf("DEBUG:GAME START %+v, %+v\n", g.bird, pipe)
		g.GameOutChan <- Message{cmd: Bird, params: []int{g.bird.xloc, g.bird.yloc}}
		g.GameOutChan <- Message{cmd: Pipe, params: []int{pipe.xloc, pipe.height}}
		log.Debug().Msgf("Starting Game: %d\n", g.game_id)
		for {
			select {
			case inMsg := <-g.GameInChan:
				//collect state changes
				{
					if inMsg.cmd == KeyPress {
						if inMsg.params[0] == 'k' {
							isKeyPressed = true
						}
					}

				}
			case <-ticker.C:
				//execute state changes
				{
					endGame := false
					log.Debug().Msgf("DEBUG:TICK\n")
					err := g.bird.UpdatePos(isKeyPressed, &g.gameOpts)
					if err != nil {
						endGame = true
					}
					status := g.Status(&g.bird, &pipe)
					g.GameOutChan <- Message{cmd: Bird, params: []int{g.bird.xloc, g.bird.yloc}}
					isKeyPressed = false
					if status == Collided {
						endGame = true
					} else if status == Crossed {
						pipe = <-g.pipe
						g.GameOutChan <- Message{cmd: Pipe, params: []int{pipe.xloc, pipe.height}}
						g.stats++
					}

					if endGame {
						g.GameOutChan <- Message{cmd: End, params: []int{g.stats}}
						g.Stop()
					}
				}
			}

		}
	}()

	return nil
}

func (g *Game) Stop() {
	//freeing resources

	g.gameOpts.ticker.Stop()
	log.Debug().Msgf("DEBUG:TICKER STOP")
    if g.ggame != nil {
        g.ggame.pipes.UnSubscribe(g.game_id)
    }

	//TODO: error from writing side
	//close(g.GameInChan)
	//close(g.GameOutChan)
}
