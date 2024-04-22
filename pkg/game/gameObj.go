package game

import (
	"fmt"
	"math/rand"
)

type Stats struct {
	time_started int
	points       int
}

type Pipe_t struct {
	xloc   int
	height int
}

func GenPipes(opts *GameOpts) chan Pipe_t {
	pipes_out := make(chan Pipe_t, 10)
	max_gap := opts.win_width / 8
	go func() {
		ppxloc := 0
		for {
			height := rand.Intn(opts.win_height/2 - 1)
			xloc := (ppxloc + max_gap/2 + rand.Intn(max_gap/2)) % opts.win_width
			ppxloc = xloc
			pipes_out <- Pipe_t{xloc: xloc, height: height}
		}
	}()
	return pipes_out
}

type Bird_t struct {
	xloc         int
	yloc         int
	vy           int
	vx           int
    jump         int
	time_started int
	opts         *GameOpts
}

func NewBird(opts *GameOpts) Bird_t {
	return Bird_t{
		xloc: 0,
		yloc: opts.win_height / 2,
		vy:   1,
		vx:   1,
        jump: 3,
		opts: opts,
	}
}

func (b *Bird_t) UpdatePos(isKeyPressed bool, gopts *GameOpts)  error {
	if isKeyPressed {
		b.yloc -= b.jump
	} else {
		b.yloc += b.vy
	}

	b.xloc = (b.xloc + b.vx) % b.opts.win_width

    if b.yloc < 0 || b.yloc > gopts.win_height {
        return fmt.Errorf("Bird out of bound")
    }

    return nil
}

