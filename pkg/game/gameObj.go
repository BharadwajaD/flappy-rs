package game

type Stats struct {
	time_started int
	points       int
}

type Pipe_t struct {
	xloc   int
	height int
}

type Bird_t struct {
	xloc         int
	yloc         int
	vy           int
	vx           int
	time_started int
}

func NewBird() Bird_t {
	return Bird_t{
		vy: 2,
		vx: 2,
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
