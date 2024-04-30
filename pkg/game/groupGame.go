package game

import (
	"math/rand"
	"github.com/bharadwajaD/flappy-go/pkg/spmc"
)

type GroupGame struct {
    group_id int  //single global group
    games []*Game //is it required ??
    pipes *spmc.Spmc[Pipe_t]
}

func NewGroupGame(opts *GameOpts) GroupGame {
    pipes := spmc.NewSpmc[Pipe_t]()

    //TODO: Needs refactoring 
	max_gap := opts.win_width / 8
    min_height := opts.win_height / 2 - 4
	go func() {
		ppxloc := 3
		for {
			height := max(min_height, rand.Intn(opts.win_height/2 - 1))
			xloc := (ppxloc + max_gap/2 + rand.Intn(max_gap/2)) % opts.win_width
			ppxloc = xloc
            //TODO: ticker might be needed
            pipes.Broadcast(Pipe_t{xloc: xloc, height: height})
		}
	}()

    return GroupGame{
        games: make([]*Game, 0, 10),
        pipes: &pipes,
    }
}

