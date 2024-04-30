package game

import "fmt"

type Bird struct {
	xloc         int
	yloc         int
	vy           int
	vx           int
	jump         int
	opts         *GameOpts
}

func NewBird(opts *GameOpts) Bird {
	return Bird{
		xloc: 0,
		yloc: opts.win_height / 2,
		vy:   1,
		vx:   1,
		jump: 3,
		opts: opts,
	}
}

// Updates pos of the bird
// error is returned if the bird goes out of bounds
func (b *Bird) UpdatePos(isKeyPressed bool, gopts *GameOpts) error {
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
