package game

import "math/rand"

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
	//max_gap := opts.win_width / 4
    vx := 1
	go func() {
		ppxloc := 0
		for {
			height := rand.Intn(opts.win_height/2 - 1)
			xloc := (ppxloc + vx) % opts.win_width
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
	time_started int
	opts         *GameOpts
}

func NewBird(opts *GameOpts) Bird_t {
	return Bird_t{
		xloc: 0,
		yloc: opts.win_height / 2,
		vy:   0,
		vx:   1,
		opts: opts,
	}
}

func (b *Bird_t) UpdatePos(isKeyPressed bool) {
	if isKeyPressed {
		b.yloc -= b.vy
	} else {
		b.yloc += b.vy
	}

	b.xloc = (b.xloc + b.vx) % b.opts.win_width
}

func (g *Game) IsCollided(p *Pipe_t, b *Bird_t) bool {
	gopts := g.gameOpts
	return (p.xloc == b.xloc) && (b.yloc <= p.height || b.yloc >= gopts.win_height-p.height)
}
