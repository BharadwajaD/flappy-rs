package game

import "time"

type GameOpts struct {
	frame_rate int
    ticker *time.Ticker
}

func NewGameOpts(fps int) GameOpts{
	ticker := time.NewTicker(time.Duration(fps) * time.Millisecond)
    return GameOpts{
        frame_rate: fps,
        ticker: ticker,
    }
}

type Game struct {
	GameChan chan Message
	Stats    string
	gameOpts GameOpts
}

func NewGame(opts GameOpts) Game {
	return Game{
		GameChan: make(chan Message),
		Stats:    "",
		gameOpts: opts,
	}
}

func (g *Game) Start() error {

    ticker := g.gameOpts.ticker

	go func() {
        for  c := range ticker.C {
			g.GameChan <- Message{obj: Pipe, param1: c.Second(), param2: 2}
			g.GameChan <- Message{obj: Bird, param1: c.Second(), param2: 18}
		}
	}()

	return nil
}

func (g *Game) Stop(){
    g.gameOpts.ticker.Stop()
}
