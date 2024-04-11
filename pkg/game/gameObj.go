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

func GenPipes(opts GameOpts) chan Pipe_t {
    pipes_out := make(chan Pipe_t, 10)
    go func (){
        for {
            height := rand.Intn(opts.win_height / 2 - 1)
            xloc := rand.Intn(opts.win_width)
            pipes_out <- Pipe_t{ xloc: xloc, height: height}
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
}

func NewBird(opts GameOpts) Bird_t {
	return Bird_t{
        xloc: 0,
        yloc: opts.win_height / 2,
		vy: 0,
		vx: 0,
	}
}

func (b *Bird_t) UpdatePos(isKeyPressed bool) {
	if isKeyPressed {
		b.yloc -= b.vy
	} else {
		b.yloc += b.vy
	}

	b.xloc += b.vx
}

func (g *Game) IsCollided(p *Pipe_t, b *Bird_t) bool {
	gopts := g.gameOpts
	return (p.xloc == b.xloc) && (b.yloc <= p.height || b.yloc >= gopts.win_height-p.height)
}
