package game

import (
	"math/rand"
)

type Pipe struct {
	xloc   int
	height int
}

// iterator for pipes
func GenPipes(opts *GameOpts) func() Pipe {

	max_gap := opts.win_width / 8
	min_height := opts.win_height/4
	ppxloc := 2

	return func() Pipe {
		height := max(min_height, rand.Intn(opts.win_height/2-1))
		xloc := (ppxloc + max_gap/2 + rand.Intn(max_gap/2)) % opts.win_width
		ppxloc = xloc
		return Pipe{xloc: xloc, height: height}
	}
}

func NewPipes(ggame *GroupGame, opts *GameOpts, game_id int) <-chan Pipe {

	if ggame != nil {
		return ggame.pipes.Subscribe(game_id)
	}

	pipes_out := make(chan Pipe, 10)
    pipes_gen := GenPipes(opts)
    go func(){
        for {
            pipe := pipes_gen()
            pipes_out <- pipe
        }
    }()
	return pipes_out
}
