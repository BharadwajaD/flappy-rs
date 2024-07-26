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

// fps, window width and height
func NewGameOpts(fps, win_width, win_height int) GameOpts {
	return GameOpts{
		frame_rate: fps,
		ticker:     time.NewTicker(time.Duration(1000/fps) * time.Millisecond),
		win_width:  win_width,
		win_height: win_height,
	}
}

func (gopts *GameOpts) Clone() GameOpts {
	//Created to get NewTicker, instead of a ptr...
	return GameOpts{
		frame_rate: gopts.frame_rate,
		ticker:     time.NewTicker(time.Duration(1000/gopts.frame_rate) * time.Millisecond),
		win_width:  gopts.win_width,
		win_height: gopts.win_height,
	}
}

type Game struct {
	game_id int
	ggame   *GroupGame

	GameOutChan chan Message
	GameInChan  chan Message
	score       int
	gameOpts    GameOpts
	bird        Bird
	pipe        <-chan Pipe
}

var game_id int

func NewGame(ggame *GroupGame, opts GameOpts) Game {

	game_id++
	return Game{
		GameOutChan: make(chan Message),
		GameInChan:  make(chan Message),
		score:       0,
		gameOpts:    opts,
		bird:        NewBird(&opts),
		pipe:        NewPipes(ggame, &opts, game_id),
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

func (g *Game) Status(b *Bird, p *Pipe) GameStatus {
	if p.xloc == b.xloc {
		if b.yloc <= p.height || b.yloc >= g.gameOpts.win_height-p.height {
			return Collided
		} else {
			return Crossed
		}
	}

	return FarAway
}

func (g *Game) Start() error {

	log.Info().Msgf("%d game started\n", g.game_id)

	ticker := g.gameOpts.ticker
	go func() {
		isKeyPressed := false
		pipe := <-g.pipe
		//starting pipe and bird
		g.GameOutChan <- Message{cmd: BirdCmd, params: []int{g.bird.xloc, g.bird.yloc}}
		g.GameOutChan <- Message{cmd: PipeCmd, params: []int{pipe.xloc, pipe.height}}

		for {
			select {
			case inMsg := <-g.GameInChan:
				//collect state changes
				{
					if inMsg.cmd == KeyPressCmd {
						if inMsg.params[0] == 'k' {
							isKeyPressed = true
						}
						if inMsg.params[0] == 'q' {
							g.GameOutChan <- Message{cmd: EndCmd, params: []int{g.score}}
							g.Stop()
						}
					}

				}
			case <-ticker.C:
				//execute state changes
				{
					log.Debug().Msgf("DEBUG:TICK\n")
					err := g.bird.UpdatePos(isKeyPressed, &g.gameOpts)
					isKeyPressed = false
					if err != nil {
						g.GameOutChan <- Message{cmd: EndCmd, params: []int{g.score}}
						g.Stop()
					}

					status := g.Status(&g.bird, &pipe)
					g.GameOutChan <- Message{cmd: BirdCmd, params: []int{g.bird.xloc, g.bird.yloc}}

					if status == Collided {
						g.GameOutChan <- Message{cmd: EndCmd, params: []int{g.score}}
						g.Stop()
					} else if status == Crossed {
						pipe = <-g.pipe
						g.GameOutChan <- Message{cmd: PipeCmd, params: []int{pipe.xloc, pipe.height}}
						g.score++
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
